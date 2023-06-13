package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// GetName get user name
//
//	@Summary		get user name by id
//	@Description	get user name by id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{string}	UserName
//	@Router			/user/get-name/{id} [get]
func (c *UserController) GetName(ctx *gin.Context, id uint) (string, error) {
	return fmt.Sprintf("User %d", id), nil
}
