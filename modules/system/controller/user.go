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
//	@Description	gget user name by idget user name by idget user name by idget user name by idet user name by id
//	@Description	gget user name by idget user name by idget user name by idget user name by idet user name by id
//	@Description	gget user name by idget user name by idget user name by idget user name by idet user name by id
//	@Description	gget user name by idget user name by idget user name by idget user name by idet user name by id
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{string}	UserName
//	@Router			/user/get-name/{id} [get]
func (c *UserController) GetName(ctx *gin.Context, id uint) (string, error) {
	return fmt.Sprintf("User %d", id), nil
}
