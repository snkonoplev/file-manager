package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

type SystemController struct {
	mediator *mediator.Mediator
}

func NewSystemController(mediator *mediator.Mediator) *SystemController {
	return &SystemController{
		mediator: mediator,
	}
}

// @Id GetDiskUsage
// @Summary Get disk usage
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/system/disk-usage [get]
// @Success 200 {object} entity.DiskUsage
// @Tags System
func (h *SystemController) GetDiskUsage(c *gin.Context) {
	result, err := h.mediator.Handle(c, query.DiskUsageQuery{})
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
