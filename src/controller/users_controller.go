package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/auth"
	"github.com/snkonoplev/file-manager/command"
	"github.com/snkonoplev/file-manager/mediator"
	"github.com/snkonoplev/file-manager/query"
)

type UsersController struct {
	mediator *mediator.Mediator
}

func NewUsersController(mediator *mediator.Mediator) *UsersController {
	return &UsersController{
		mediator: mediator,
	}
}

// @Id GetUsers
// @Summary Get list of all users
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/users [get]
// @Success 200 {object} []entity.User
// @Failure 401 {string} string
// @Tags Users
func (h *UsersController) GetUsers(c *gin.Context) {
	result, err := h.mediator.Handle(c, query.UsersQuery{})
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id CreteUser
// @Summary Crete new user
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param Body body command.CreateUserCommand true "User"
// @Router /api/users [post]
// @Success 200 {object} int
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Tags Users
func (h *UsersController) CreteUser(c *gin.Context) {

	var user command.CreateUserCommand
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, "can't bind user model")
		return
	}

	if claims, ok := c.Get(auth.Claims); ok {
		claim := claims.(*auth.Claim)
		user.IsCallerAdmin = claim.IsAdmin
	}

	result, err := h.mediator.Handle(c, user)
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id UpdateUser
// @Summary Update user
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param Body body command.UpdateUserCommand true "User"
// @Router /api/users [put]
// @Success 200 {object} int
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Tags Users
func (h *UsersController) UpdateUser(c *gin.Context) {
	var user command.UpdateUserCommand
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, "can't bind user model")
		return
	}

	if claims, ok := c.Get(auth.Claims); ok {
		claim := claims.(*auth.Claim)
		user.IsCallerAdmin = claim.IsAdmin
	}

	result, err := h.mediator.Handle(c, user)
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id DeleteUser
// @Summary Delete user
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path int true "User id"
// @Router /api/users/{id} [delete]
// @Success 200 {object} int
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Tags Users
func (h *UsersController) DeleteUser(c *gin.Context) {

	strId := c.Param("id")

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("can't convert %s to int", strId))
		return
	}

	user := command.DeleteUserCommand{
		Id: id,
	}

	if claims, ok := c.Get(auth.Claims); ok {
		claim := claims.(*auth.Claim)
		user.IsCallerAdmin = claim.IsAdmin
	}

	result, err := h.mediator.Handle(c, user)
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id GetUser
// @Summary Get user by id
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path int true "User id"
// @Router /api/users/{id} [get]
// @Success 200 {object} entity.User
// @Failure 401 {string} string
// @Tags Users
func (h *UsersController) GetUser(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("can't convert %s to int", strId))
		return
	}

	q := query.UserQuery{
		Id: id,
	}

	result, err := h.mediator.Handle(c, q)
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id CurrentUser
// @Summary Get current
// @Accept  json
// @Produce  json
// @Security Bearer
// @Router /api/users/current [get]
// @Success 200 {object} entity.User
// @Failure 401 {string} string
// @Tags Users
func (h *UsersController) CurrentUser(c *gin.Context) {

	q := query.UserQuery{}

	if claims, ok := c.Get(auth.Claims); ok {
		claim := claims.(*auth.Claim)
		q.Id = claim.UserId
	} else {
		c.String(http.StatusBadRequest, "can't get user id from claims")
	}

	result, err := h.mediator.Handle(c, q)
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Id ChangePassword
// @Summary Change password
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param Body body command.ChangePasswordCommand true "Password"
// @Router /api/users/change-password [put]
// @Success 200 {object} int
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Tags Users
func (h *UsersController) ChangePassword(c *gin.Context) {

	var passwordCommand command.ChangePasswordCommand
	if err := c.ShouldBind(&passwordCommand); err != nil {
		c.String(http.StatusBadRequest, "can't bind password model")
		return
	}

	if claims, ok := c.Get(auth.Claims); ok {
		claim := claims.(*auth.Claim)
		passwordCommand.IsCallerAdmin = claim.IsAdmin
	}

	result, err := h.mediator.Handle(c, passwordCommand)
	if err != nil {
		target := &mediator.HandlerError{}
		if errors.As(err, &target) {
			c.String(target.StatusCode, target.Message)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}
