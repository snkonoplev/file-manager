package entity

type CpuUsage struct {
	CountLogical  int       `json:"countLogical" example:"1"`
	CountPhysical int       `json:"countPhysical" example:"1"`
	Percent       []float64 `json:"percent" example:"0.001"`
}
