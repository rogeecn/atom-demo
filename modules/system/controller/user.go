package controller

import (
	"atom/http/modules/system/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// GetName get user name
//
//	@Summary		get user by id
//	@Description	get user info by id
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/users/{id} [get]
func (c *UserController) Show(ctx *gin.Context, id uint) (*dto.UserItem, error) {
	return &dto.UserItem{
		Name: "Rogee",
	}, nil
}

// List list user by query filter
//
//	@Summary		list user by query filter
//	@Description	list user by query filter
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query	dto.UserListQueryFilter	true	"QueryFilter"
//	@Success		200			{array}	dto.UserItem
//	@Router			/users [get]
func (c *UserController) List(ctx *gin.Context, queryFilter *dto.UserListQueryFilter) ([]*dto.UserItem, error) {
	return []*dto.UserItem{
		{
			Name: "Rogee",
		},
		{
			Name: "Rogee",
		},
	}, nil
}

// Create create a new user
//
//	@Summary		create new user
//	@Description	create new user
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.UserForm	true	"UserForm"
//	@Success		200		{int}	UserID
//	@Router			/users [post]
func (c *UserController) Create(ctx *gin.Context, body *dto.UserForm) (int, error) {
	return 12, nil
}

// Update update user by id
//
//	@Summary		update user by id
//	@Description	update user by id
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"UserID"
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Failure		500		{string}	UserID
//	@Router			/users/{id} [put]
func (c *UserController) Update(ctx *gin.Context, id uint, body *dto.UserForm) error {
	return nil
}

// Delete delete user by id
//
//	@Summary		delete user by id
//	@Description	delete user by id
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{string}	UserID
//	@Failure		500	{string}	UserID
//	@Router			/users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context, id uint) error {
	return nil
}
