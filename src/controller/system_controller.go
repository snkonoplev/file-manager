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

// @Id GetMemoryUsage
// @Summary Get memory usage
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/system/memory-usage [get]
// @Success 200 {object} entity.MemoryUsage
// @Tags System
func (h *SystemController) GetMemoryUsage(c *gin.Context) {
	result, err := h.mediator.Handle(c, query.MemoryUsageQuery{})
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

// @Id GetCpuUsage
// @Summary Get CPU usage
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/system/cpu-usage [get]
// @Success 200 {object} entity.CpuUsage
// @Tags System
func (h *SystemController) GetCpuUsage(c *gin.Context) {
	result, err := h.mediator.Handle(c, query.CpuUsageQuery{})
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

// @Id GetLoadAvg
// @Summary Get Get load avg
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/system/load-avg [get]
// @Success 200 {object} entity.LoadAvg
// @Tags System
func (h *SystemController) GetLoadAvg(c *gin.Context) {
	result, err := h.mediator.Handle(c, query.LoadAvgQuery{})
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

// @Id GetUpTime
// @Summary Get Get up time
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/system/up-time [get]
// @Success 200 {object} entity.UpTime
// @Tags System
func (h *SystemController) GetUpTime(c *gin.Context) {
	result, err := h.mediator.Handle(c, query.UpTimeQuery{})
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
