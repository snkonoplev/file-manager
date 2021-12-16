package queryhandler

import (
	"context"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/snkonoplev/file-manager/entity"
)

type CpuUsageHandler struct {
}

func NewCpuUsageHandler() *CpuUsageHandler {
	return &CpuUsageHandler{}
}

func (h *CpuUsageHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	cpu, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	return entity.CpuUsage{
		Count:  cpu.CPUCount,
		Total:  cpu.Total,
		User:   cpu.User,
		System: cpu.System,
		Idle:   cpu.Idle,
	}, nil
}
