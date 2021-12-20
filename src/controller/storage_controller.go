package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

type StorageController struct {
	mediator *mediator.Mediator
}

func NewStorageController(mediator *mediator.Mediator) *StorageController {
	return &StorageController{
		mediator: mediator,
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
