package entity

type CpuUsage struct {
	Count  int     `json:"count" example:"1637768672"`
	Total  float64 `json:"total" example:"1637768672"`
	User   float64 `json:"user" example:"1637768672"`
	System float64 `json:"system" example:"1637768672"`
	Idle   float64 `json:"idle" example:"1637768672"`
}
