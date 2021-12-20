package controller

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"

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
// @Param path query string false "Path"
// @Router /api/storage [get]
// @Success 200 {object} []entity.DirectoryDataWrapper
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Tags Storage
func (h *StorageController) GetDirectoryContent(c *gin.Context) {
	path := c.DefaultQuery("path", ".")
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
// @Param file query string false "File"
// @Router /api/storage/download [get]
// @Success 200 {object} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Tags Storage
func (h *StorageController) DownloadFile(c *gin.Context) {
	filePath := c.DefaultQuery("file", ".")
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
