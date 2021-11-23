package query

type UserAuthorizeQuery struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
