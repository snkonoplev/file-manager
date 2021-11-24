package command

type CreateUserCommand struct {
	Name          string `db:"name" json:"name" example:"adam"`
	Password      string `db:"password" json:"password" example:"123"`
	IsAdmin       bool   `db:"isAdmin" json:"is_admin" example:"false"`
	IsCallerAdmin bool   `db:"-" json:"-"`
}
