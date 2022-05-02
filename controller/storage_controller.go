package controller

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
	"github.com/spf13/viper"
)

type StorageController struct {
	mediator    *mediator.Mediator
	storagePath string
}

func NewStorageController(mediator *mediator.Mediator, viper *viper.Viper) *StorageController {
	return &StorageController{
		mediator:    mediator,
		storagePath: viper.GetString("STORAGE_PATH"),
	}
}

// @Id GetDirectoryContent
// @Summary Get directory content
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param directory path string true "Directory"
// @Router /api/storage/list-directories/{directory} [get]
// @Success 200 {object} []entity.DirectoryDataWrapper
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Tags Storage
func (h *StorageController) GetDirectoryContent(c *gin.Context) {

	path := c.Param("directory")

	result, err := h.mediator.Handle(c, query.ReadDirectoryQuery{Path: path})
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id DownloadFile
// @Summary Download file from storage
// @Accept  json
// @Produce  octet-stream
// @Security Bearer
// @Param file path string true "File"
// @Router /api/storage/download/{file} [get]
// @Success 200 {object} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Tags Storage
func (h *StorageController) DownloadFile(c *gin.Context) {
	filePath := c.Param("file")
	file, err := os.Open(path.Join(h.storagePath, filePath))
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-type", "application/octet-stream")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
}

// @Id UploadFile
// @Summary Upload File
// @Accept  mpfd
// @Produce json
// @Security Bearer
// @Router /api/upload [post]
// @Success 200 {object} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Tags Storage
func (h *StorageController) UploadFile(c *gin.Context) {

	formdata, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "can't read MultipartForm")
		return
	}

	filePath := formdata.Value["path"]
	if len(filePath) < 1 {
		c.String(http.StatusBadRequest, "can't read path")
		return
	}

	var folder string

	if filePath[0] == "." {
		folder = h.storagePath
	} else {
		folder = path.Join(h.storagePath, filePath[0])
	}

	if _, err = os.Stat(folder); os.IsNotExist(err) {
		c.String(http.StatusBadRequest, "selected directory is not exists")
		return
	}

	for _, fh := range formdata.File {
		for _, fileHeader := range fh {
			file, err := fileHeader.Open()
			if err != nil {
				c.String(http.StatusBadRequest, "can't open file header")
				return
			}
			defer file.Close()

			f, err := os.Create(filepath.Join(folder, fileHeader.Filename))
			if err != nil {
				c.String(http.StatusBadRequest, "can't create file")
				return
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				c.String(http.StatusBadRequest, "can't save file")
				return
			}
		}
	}

	c.Status(http.StatusOK)
}

// @Id DeleteFile
// @Summary Remove file from storage
// @Accept  json
// @Produce json
// @Security Bearer
// @Param file path string true "File"
// @Router /api/storage/delete/{file} [delete]
// @Success 200 {object} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Tags Storage
func (h *StorageController) DeleteFile(c *gin.Context) {
	filePath := c.Param("file")
	err := os.Remove(path.Join(h.storagePath, filePath))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

// @Id MkDir
// @Summary Create directory in storage
// @Accept  json
// @Produce json
// @Security Bearer
// @Param dir path string true "Directory"
// @Router /api/storage/create-directory/{dir} [put]
// @Success 200 {object} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Tags Storage
func (h *StorageController) MkDir(c *gin.Context) {
	dirPath := path.Join(h.storagePath, c.Param("dir"))
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	} else {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}
