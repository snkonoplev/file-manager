package entity

type LoadAvg struct {
	Loadavg1  float64 `json:"loadavg1" example:"1637768672"`
	Loadavg5  float64 `json:"loadavg5" example:"1637768672"`
	Loadavg15 float64 `json:"loadavg15" example:"1637768672"`
}
