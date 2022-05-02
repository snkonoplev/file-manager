package queryhandler

import (
	"context"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/snkonoplev/file-manager/entity"
	"github.com/spf13/viper"
)

type DiskUsageHandler struct {
	volumePath string
}

func NewDiskUsageHandler(viper *viper.Viper) *DiskUsageHandler {
	return &DiskUsageHandler{
		volumePath: viper.GetString("STORAGE_PATH"),
	}
}

func (h *DiskUsageHandler) Handle(context context.Context, q interface{}) (interface{}, error) {
	usage := du.NewDiskUsage(h.volumePath)
	return entity.DiskUsage{
		Available: usage.Available(),
		Size:      usage.Size(),
		Used:      usage.Used(),
		Usage:     usage.Usage(),
	}, nil
}
