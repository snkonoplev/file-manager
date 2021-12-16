package queryhandler

import (
	"context"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/snkonoplev/file-manager/entity"
)

type CpuUsageHandler struct {
}

func NewCpuUsageHandler() *CpuUsageHandler {
	return &CpuUsageHandler{}
}

func (h *CpuUsageHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	before, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	return entity.CpuUsage{
		Count:  after.CPUCount,
		Total:  float64(after.Total - before.Total),
		User:   float64(after.User - before.User),
		System: float64(after.System - before.System),
		Idle:   float64(after.Idle - before.Idle),
	}, nil
}
