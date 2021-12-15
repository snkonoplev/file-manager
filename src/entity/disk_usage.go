package entity

type DiskUsage struct {
	Available uint64  `json:"available" example:"1637768672"`
	Size      uint64  `json:"size" example:"1637768672"`
	Used      uint64  `json:"used" example:"1637768672"`
	Usage     float32 `json:"usage" example:"1637768672"`
}
