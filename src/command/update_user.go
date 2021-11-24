package command

type UpdateUserCommand struct {
	Id            int64  `db:"id" json:"id" example:"1"`
	Name          string `db:"name" json:"name" example:"Adam"`
	IsAdmin       bool   `db:"is_admin" json:"isAdmin" example:"true"`
	Password      string `db:"password" json:"password"`
	IsCallerAdmin bool   `db:"-" json:"-"`
}
