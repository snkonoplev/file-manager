package entity

type User struct {
	Id        int64  `db:"id" json:"id" example:"1"`
	Created   int64  `db:"created" json:"created" example:"1637768672"`
	LastLogin *int64 `db:"last_login" json:"lastLogin" example:"1637768672"`
	Name      string `db:"name" json:"name" example:"Adam"`
	IsAdmin   bool   `db:"is_admin" json:"isAdmin" example:"true"`
	Password  string `db:"password" json:"-"`
}
