package command

type ChangePasswordCommand struct {
	Name             string `json:"name" example:"adam" binding:"required"`
	NewPassword      string `json:"newPassword" example:"123" binding:"required"`
	PreviousPassword string `json:"previousPassword" example:"123" binding:"required"`
	ClaimUserId      int64  `json:"-"`
}
