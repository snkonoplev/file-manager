package queryhandler

import (
	"context"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/snkonoplev/file-manager/entity"
)

type MemoryUsageHandler struct {
}

func NewMemoryUsageHandler() *MemoryUsageHandler {
	return &MemoryUsageHandler{}
}

func (h *MemoryUsageHandler) Handle(context context.Context, q interface{}) (interface{}, error) {

	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return entity.MemoryUsage{
		Total:     v.Total,
		Used:      v.Used,
		Cached:    v.Cached,
		Free:      v.Free,
		Available: v.Available,
	}, nil
}
