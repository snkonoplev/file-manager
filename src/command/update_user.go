package command

type UpdateUserCommand struct {
	Id            int64 `db:"id" json:"id" example:"1"`
	IsAdmin       bool  `db:"is_admin" json:"isAdmin" example:"true"`
	IsCallerAdmin bool  `db:"-" json:"-"`
}
