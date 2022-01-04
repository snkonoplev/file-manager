package entity

type DirectoryData struct {
	Name string `json:"name" example:"test"`
	Size int64  `json:"size" example:"10"`
	Type string `json:"type" example:"File"`
}
