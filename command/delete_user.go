package command

type DeleteUserCommand struct {
	Id            int64 `json:"id" example:"1" binding:"required"`
	IsCallerAdmin bool  `db:"-" json:"-"`
}
