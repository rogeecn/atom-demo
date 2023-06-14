package service

import (
	"context"

	"atom/http/common"
	"atom/http/database/models"
	"atom/http/modules/system/dao"
	"atom/http/modules/system/dto"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(
	userDao *dao.UserDao,
) *UserService {
	return &UserService{
		userDao: userDao,
	}
}

func (svc *UserService) GetByID(ctx context.Context, id int32) (*dto.UserItem, error) {
	userModel, err := svc.userDao.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get user by id failed")
	}

	resp := &dto.UserItem{}
	_ = copier.Copy(resp, userModel)

	return resp, nil
}

func (svc *UserService) PageByQueryFilter(ctx context.Context, pageFilter *common.PageQueryFilter, queryFilter *dto.UserListQueryFilter) ([]*dto.UserItem, int64, error) {
	userModels, total, err := svc.userDao.PageByQueryFilter(ctx, pageFilter.Format(), queryFilter)
	if err != nil {
		return nil, 0, err
	}

	resp := []*dto.UserItem{}
	for _, u := range userModels {
		item := &dto.UserItem{}
		_ = copier.Copy(item, u)
		resp = append(resp, item)
	}

	return resp, total, nil
}

// Create
func (svc *UserService) Create(ctx context.Context, body *dto.UserForm) error {
	model := &models.User{}
	_ = copier.Copy(model, body)
	return svc.userDao.Create(ctx, model)
}

// Update
func (svc *UserService) Update(ctx context.Context, id int32, body *dto.UserForm) error {
	model := &models.User{}
	_ = copier.Copy(model, body)
	model.ID = id
	return svc.userDao.Update(ctx, id, model)
}

// Delete
func (svc *UserService) Delete(ctx context.Context, id int32) error {
	return svc.userDao.Delete(ctx, id)
}
