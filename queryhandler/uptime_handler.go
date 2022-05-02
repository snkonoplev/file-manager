package queryhandler

import (
	"context"

	"github.com/mackerelio/go-osstat/uptime"
	"github.com/snkonoplev/file-manager/entity"
)

type UpTimeHandler struct {
}

func NewUpTimeHandler() *UpTimeHandler {
	return &UpTimeHandler{}
}

func (h *UpTimeHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	uptime, err := uptime.Get()
	if err != nil {
		return nil, err
	}

	return entity.UpTime{
		UpTimeMilliseconds: uptime.Milliseconds(),
	}, nil
}
