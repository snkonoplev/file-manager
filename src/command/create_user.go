package command

type CreateUserCommand struct {
	Name          string `db:"name" json:"name" example:"adam" binding:"required"`
	Password      string `db:"password" json:"password" example:"123" binding:"required"`
	IsAdmin       bool   `db:"is_admin" json:"isAdmin" example:"false"`
	IsActive      bool   `db:"is_active" json:"isActive" example:"true"`
	IsCallerAdmin bool   `db:"-" json:"-"`
}
