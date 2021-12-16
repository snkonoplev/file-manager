package entity

type MemoryUsage struct {
	Total     uint64 `json:"total" example:"1637768672"`
	Used      uint64 `json:"used" example:"1637768672"`
	Cached    uint64 `json:"cached" example:"1637768672"`
	Free      uint64 `json:"free" example:"1637768672"`
	Available uint64 `json:"available" example:"1637768672"`
}
