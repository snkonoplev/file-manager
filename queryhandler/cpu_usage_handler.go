package queryhandler

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"

	"github.com/snkonoplev/file-manager/entity"
)

type CpuUsageHandler struct {
}

func NewCpuUsageHandler() *CpuUsageHandler {
	return &CpuUsageHandler{}
}

func (h *CpuUsageHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	countL, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	countP, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}

	percent, err := cpu.PercentWithContext(context, time.Second, false)
	if err != nil {
		return nil, err
	}

	return entity.CpuUsage{
		Percent:       percent,
		CountLogical:  countL,
		CountPhysical: countP,
	}, nil
}
