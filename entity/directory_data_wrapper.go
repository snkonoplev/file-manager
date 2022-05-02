package entity

type DirectoryDataWrapper struct {
	Key  string        `json:"key" example:"test/test2"`
	Data DirectoryData `json:"data"`
	Leaf bool          `json:"leaf"`
}
