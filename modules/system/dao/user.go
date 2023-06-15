package dao

import (
	"atom/http/common"
	"atom/http/database/models"
	"atom/http/database/query"
	"atom/http/modules/system/dto"
	"context"

	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
)

type UserDao struct {
	query *query.Query
}

func NewUserDao(query *query.Query) *UserDao {
	return &UserDao{query: query}
}

func (dao *UserDao) Update(ctx context.Context, id int32, model *models.User) error {
	oldModel, err := dao.GetByID(ctx, id)
	if err != nil {
		return err
	}
	_ = copier.Copy(oldModel, model)

	user := dao.query.User
	_, err = user.WithContext(ctx).Where(user.ID.Eq(id)).Updates(model)

	return err
}

func (dao *UserDao) Delete(ctx context.Context, id int32) error {
	user := dao.query.User
	_, err := user.WithContext(ctx).Where(user.ID.Eq(id)).Delete()
	return err
}

func (dao *UserDao) Create(ctx context.Context, model *models.User) error {
	return dao.query.User.WithContext(ctx).Create(model)
}

func (dao *UserDao) GetByID(ctx context.Context, id int32) (*models.User, error) {
	user := dao.query.User
	return user.WithContext(ctx).Where(user.ID.Eq(id)).First()
}

func (dao *UserDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, int64, error) {
	user := dao.query.User
	userQuery := user.WithContext(ctx)
	if queryFilter != nil {
		if queryFilter.ID != nil {
			userQuery = userQuery.Where(user.ID.Eq(*queryFilter.ID))
		}

		if queryFilter.Username != nil {
			userQuery = userQuery.Where(user.Username.Like(query.WrapLike(*queryFilter.Username)))
		}

		if queryFilter.Age != nil {
			userQuery = userQuery.Where(user.Age.Eq(*queryFilter.Age))
		}

		if queryFilter.Sex != nil {
			userQuery = userQuery.Where(user.Sex.Eq(*queryFilter.Sex))
		}

		if queryFilter.Birthday != nil {
			userQuery = userQuery.Where(user.Birthday.Eq(*queryFilter.Birthday))
		}

		if queryFilter.Status != nil {
			userQuery = userQuery.Where(user.Status.Eq(*queryFilter.Status))
		}

		if queryFilter.State != nil {
			userQuery = userQuery.Where(user.State.Eq(*queryFilter.State))
		}
	}

	if sortFilter != nil {
		orderExprs := []field.Expr{}
		for _, v := range sortFilter.Asc {
			if expr, ok := user.GetFieldByName(v); ok {
				orderExprs = append(orderExprs, expr)
			}
		}
		for _, v := range sortFilter.Desc {
			if expr, ok := user.GetFieldByName(v); ok {
				orderExprs = append(orderExprs, expr.Desc())
			}
		}
		userQuery = userQuery.Order(orderExprs...)
	}

	return userQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *UserDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, error) {
	user := dao.query.User
	userQuery := user.WithContext(ctx)
	if queryFilter != nil {
		if queryFilter.ID != nil {
			userQuery = userQuery.Where(user.ID.Eq(*queryFilter.ID))
		}

		if queryFilter.Username != nil {
			userQuery = userQuery.Where(user.Username.Like(query.WrapLike(*queryFilter.Username)))
		}

		if queryFilter.Age != nil {
			userQuery = userQuery.Where(user.Age.Eq(*queryFilter.Age))
		}

		if queryFilter.Sex != nil {
			userQuery = userQuery.Where(user.Sex.Eq(*queryFilter.Sex))
		}

		if queryFilter.Birthday != nil {
			userQuery = userQuery.Where(user.Birthday.Eq(*queryFilter.Birthday))
		}

		if queryFilter.Status != nil {
			userQuery = userQuery.Where(user.Status.Eq(*queryFilter.Status))
		}

		if queryFilter.State != nil {
			userQuery = userQuery.Where(user.State.Eq(*queryFilter.State))
		}
	}

	if sortFilter != nil {
		orderExprs := []field.Expr{}
		for _, v := range sortFilter.Asc {
			if expr, ok := user.GetFieldByName(v); ok {
				orderExprs = append(orderExprs, expr)
			}
		}
		for _, v := range sortFilter.Desc {
			if expr, ok := user.GetFieldByName(v); ok {
				orderExprs = append(orderExprs, expr.Desc())
			}
		}
		userQuery = userQuery.Order(orderExprs...)
	}

	return userQuery.Find()
}
