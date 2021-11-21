package entity

type User struct {
	Id        int64  `db:"id" json:"id" example:"1"`
	Created   int64  `db:"created" json:"created" example:"1"`
	LastLogin *int64 `db:"last_login" json:"lastLogin" example:"1"`
	Name      string `db:"name" json:"name" example:"1"`
	IsAdmin   int64  `db:"is_admin" json:"isAdmin" example:"1"`
}
