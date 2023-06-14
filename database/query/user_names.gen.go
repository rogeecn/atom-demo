// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"atom/http/database/models"
)

func newUserName(db *gorm.DB, opts ...gen.DOOption) userName {
	_userName := userName{}

	_userName.userNameDo.UseDB(db, opts...)
	_userName.userNameDo.UseModel(&models.UserName{})

	tableName := _userName.userNameDo.TableName()
	_userName.ALL = field.NewAsterisk(tableName)
	_userName.UserName = field.NewString(tableName, "user_name")
	_userName.Password = field.NewString(tableName, "password")

	_userName.fillFieldMap()

	return _userName
}

type userName struct {
	userNameDo userNameDo

	ALL      field.Asterisk
	UserName field.String
	Password field.String

	fieldMap map[string]field.Expr
}

func (u userName) Table(newTableName string) *userName {
	u.userNameDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userName) As(alias string) *userName {
	u.userNameDo.DO = *(u.userNameDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userName) updateTableName(table string) *userName {
	u.ALL = field.NewAsterisk(table)
	u.UserName = field.NewString(table, "user_name")
	u.Password = field.NewString(table, "password")

	u.fillFieldMap()

	return u
}

func (u *userName) WithContext(ctx context.Context) IUserNameDo { return u.userNameDo.WithContext(ctx) }

func (u userName) TableName() string { return u.userNameDo.TableName() }

func (u userName) Alias() string { return u.userNameDo.Alias() }

func (u *userName) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userName) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["user_name"] = u.UserName
	u.fieldMap["password"] = u.Password
}

func (u userName) clone(db *gorm.DB) userName {
	u.userNameDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userName) replaceDB(db *gorm.DB) userName {
	u.userNameDo.ReplaceDB(db)
	return u
}

type userNameDo struct{ gen.DO }

type IUserNameDo interface {
	gen.SubQuery
	Debug() IUserNameDo
	WithContext(ctx context.Context) IUserNameDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserNameDo
	WriteDB() IUserNameDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserNameDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserNameDo
	Not(conds ...gen.Condition) IUserNameDo
	Or(conds ...gen.Condition) IUserNameDo
	Select(conds ...field.Expr) IUserNameDo
	Where(conds ...gen.Condition) IUserNameDo
	Order(conds ...field.Expr) IUserNameDo
	Distinct(cols ...field.Expr) IUserNameDo
	Omit(cols ...field.Expr) IUserNameDo
	Join(table schema.Tabler, on ...field.Expr) IUserNameDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserNameDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserNameDo
	Group(cols ...field.Expr) IUserNameDo
	Having(conds ...gen.Condition) IUserNameDo
	Limit(limit int) IUserNameDo
	Offset(offset int) IUserNameDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserNameDo
	Unscoped() IUserNameDo
	Create(values ...*models.UserName) error
	CreateInBatches(values []*models.UserName, batchSize int) error
	Save(values ...*models.UserName) error
	First() (*models.UserName, error)
	Take() (*models.UserName, error)
	Last() (*models.UserName, error)
	Find() ([]*models.UserName, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UserName, err error)
	FindInBatches(result *[]*models.UserName, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.UserName) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserNameDo
	Assign(attrs ...field.AssignExpr) IUserNameDo
	Joins(fields ...field.RelationField) IUserNameDo
	Preload(fields ...field.RelationField) IUserNameDo
	FirstOrInit() (*models.UserName, error)
	FirstOrCreate() (*models.UserName, error)
	FindByPage(offset int, limit int) (result []*models.UserName, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserNameDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userNameDo) Debug() IUserNameDo {
	return u.withDO(u.DO.Debug())
}

func (u userNameDo) WithContext(ctx context.Context) IUserNameDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userNameDo) ReadDB() IUserNameDo {
	return u.Clauses(dbresolver.Read)
}

func (u userNameDo) WriteDB() IUserNameDo {
	return u.Clauses(dbresolver.Write)
}

func (u userNameDo) Session(config *gorm.Session) IUserNameDo {
	return u.withDO(u.DO.Session(config))
}

func (u userNameDo) Clauses(conds ...clause.Expression) IUserNameDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userNameDo) Returning(value interface{}, columns ...string) IUserNameDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userNameDo) Not(conds ...gen.Condition) IUserNameDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userNameDo) Or(conds ...gen.Condition) IUserNameDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userNameDo) Select(conds ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userNameDo) Where(conds ...gen.Condition) IUserNameDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userNameDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserNameDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userNameDo) Order(conds ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userNameDo) Distinct(cols ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userNameDo) Omit(cols ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userNameDo) Join(table schema.Tabler, on ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userNameDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userNameDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userNameDo) Group(cols ...field.Expr) IUserNameDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userNameDo) Having(conds ...gen.Condition) IUserNameDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userNameDo) Limit(limit int) IUserNameDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userNameDo) Offset(offset int) IUserNameDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userNameDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserNameDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userNameDo) Unscoped() IUserNameDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userNameDo) Create(values ...*models.UserName) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userNameDo) CreateInBatches(values []*models.UserName, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userNameDo) Save(values ...*models.UserName) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userNameDo) First() (*models.UserName, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserName), nil
	}
}

func (u userNameDo) Take() (*models.UserName, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserName), nil
	}
}

func (u userNameDo) Last() (*models.UserName, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserName), nil
	}
}

func (u userNameDo) Find() ([]*models.UserName, error) {
	result, err := u.DO.Find()
	return result.([]*models.UserName), err
}

func (u userNameDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UserName, err error) {
	buf := make([]*models.UserName, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userNameDo) FindInBatches(result *[]*models.UserName, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userNameDo) Attrs(attrs ...field.AssignExpr) IUserNameDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userNameDo) Assign(attrs ...field.AssignExpr) IUserNameDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userNameDo) Joins(fields ...field.RelationField) IUserNameDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userNameDo) Preload(fields ...field.RelationField) IUserNameDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userNameDo) FirstOrInit() (*models.UserName, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserName), nil
	}
}

func (u userNameDo) FirstOrCreate() (*models.UserName, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserName), nil
	}
}

func (u userNameDo) FindByPage(offset int, limit int) (result []*models.UserName, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userNameDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userNameDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userNameDo) Delete(models ...*models.UserName) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userNameDo) withDO(do gen.Dao) *userNameDo {
	u.DO = *do.(*gen.DO)
	return u
}