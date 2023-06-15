package controller

import (
	"atom/http/common"
	"atom/http/modules/system/dto"
	"atom/http/modules/system/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc *service.UserService
}

func NewUserController(
	userSvc *service.UserService,
) *UserController {
	return &UserController{
		userSvc: userSvc,
	}
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
func (c *UserController) Show(ctx *gin.Context, id int32) (*dto.UserItem, error) {
	return c.userSvc.GetByID(ctx, id)
}

// List list user by query filter
//
//	@Summary		list user by query filter
//	@Description	list user by query filter
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse
//	@Router			/users [get]
func (c *UserController) List(
	ctx *gin.Context,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.userSvc.PageByQueryFilter(ctx, queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           items,
	}, nil
}

// Create create a new user
//
//	@Summary		create new user
//	@Description	create new user
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Router			/users [post]
func (c *UserController) Create(ctx *gin.Context, body *dto.UserForm) error {
	return c.userSvc.Create(ctx, body)
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
func (c *UserController) Update(ctx *gin.Context, id int32, body *dto.UserForm) error {
	return c.userSvc.Update(ctx, id, body)
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
func (c *UserController) Delete(ctx *gin.Context, id int32) error {
	return c.userSvc.Delete(ctx, id)
}
