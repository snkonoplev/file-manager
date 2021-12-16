package queryhandler

import (
	"context"

	"github.com/mackerelio/go-osstat/memory"
	"github.com/snkonoplev/file-manager/entity"
)

type MemoryUsageHandler struct {
}

func NewMemoryUsageHandler() *MemoryUsageHandler {
	return &MemoryUsageHandler{}
}

func (h *MemoryUsageHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	memory, err := memory.Get()
	if err != nil {
		return nil, err
	}

	return entity.MemoryUsage{
		Total:     memory.Total,
		Used:      memory.Used,
		Cached:    memory.Cached,
		Free:      memory.Free,
		Available: memory.Available,
	}, nil
}
