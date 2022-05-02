package query

type UserAuthorizeQuery struct {
	UserName string `form:"username" json:"username" binding:"required" example:"admin"`
	Password string `form:"password" json:"password" binding:"required" example:"123"`
}
