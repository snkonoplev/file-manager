package command

type UpdateUserCommand struct {
	Id            int64 `db:"id" json:"id" example:"1" binding:"required"`
	IsAdmin       bool  `db:"is_admin" json:"isAdmin" example:"false"`
	IsActive      bool  `db:"is_active" json:"isActive" example:"false"`
	IsCallerAdmin bool  `db:"-" json:"-"`
}
