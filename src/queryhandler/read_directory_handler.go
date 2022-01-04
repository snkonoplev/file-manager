package queryhandler

import (
	"context"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/snkonoplev/file-manager/entity"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
	"github.com/spf13/viper"
)

type ReadDirectoryHandler struct {
	BasePath string
}

func NewReadDirectoryHandler(viper *viper.Viper) *ReadDirectoryHandler {
	return &ReadDirectoryHandler{
		BasePath: viper.GetString("STORAGE_PATH"),
	}
}

func (h *ReadDirectoryHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	if query, ok := q.(query.ReadDirectoryQuery); ok {
		p := h.BasePath
		if query.Path != "." {
			p = path.Join(p, query.Path)
		}

		files, err := ioutil.ReadDir(p)
		if err != nil {
			return nil, err
		}

		data := []entity.DirectoryDataWrapper{}

		for _, f := range files {

			key := f.Name()

			if query.Path != "." {
				key = path.Join(query.Path, f.Name())
			}

			wrapper := entity.DirectoryDataWrapper{
				Key:  key,
				Data: entity.DirectoryData{},
				Leaf: true,
			}

			if f.IsDir() {
				wrapper.Data.Type = "Folder"
				d, err := ioutil.ReadDir(path.Join(h.BasePath, key))
				if err != nil {
					return nil, err
				}
				if len(d) > 0 {
					wrapper.Leaf = false
				}
			} else {
				wrapper.Data.Type = "File"
			}

			wrapper.Data.Name = f.Name()
			wrapper.Data.Size = f.Size()

			data = append(data, wrapper)
		}
		return data, nil
	}

	return nil, &mediator.HandlerError{
		StatusCode: http.StatusInternalServerError,
		Message:    "wrong command type",
	}
}
