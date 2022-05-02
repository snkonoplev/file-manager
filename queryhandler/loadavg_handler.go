package queryhandler

import (
	"context"

	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/snkonoplev/file-manager/entity"
)

type LoadAvgHandler struct {
}

func NewLoadAvgHandler() *LoadAvgHandler {
	return &LoadAvgHandler{}
}

func (h *LoadAvgHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	loadavg, err := loadavg.Get()
	if err != nil {
		return nil, err
	}

	return entity.LoadAvg{
		Loadavg1:  loadavg.Loadavg1,
		Loadavg5:  loadavg.Loadavg5,
		Loadavg15: loadavg.Loadavg15,
	}, nil
}
