package entity

type CpuUsage struct {
	Count  int    `json:"count" example:"1637768672"`
	Total  uint64 `json:"total" example:"1637768672"`
	User   uint64 `json:"user" example:"1637768672"`
	System uint64 `json:"system" example:"1637768672"`
	Idle   uint64 `json:"idle" example:"1637768672"`
}
