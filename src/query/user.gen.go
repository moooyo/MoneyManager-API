// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"MoneyManager/src/model"
)

func newData(db *gorm.DB, opts ...gen.DOOption) data {
	_data := data{}

	_data.dataDo.UseDB(db, opts...)
	_data.dataDo.UseModel(&model.Data{})

	tableName := _data.dataDo.TableName()
	_data.ALL = field.NewAsterisk(tableName)
	_data.ID = field.NewUint(tableName, "id")
	_data.UserName = field.NewString(tableName, "username")
	_data.Password = field.NewString(tableName, "password")

	_data.fillFieldMap()

	return _data
}

type data struct {
	dataDo dataDo

	ALL      field.Asterisk
	ID       field.Uint   // key id
	UserName field.String // username
	Password field.String // password

	fieldMap map[string]field.Expr
}

func (d data) Table(newTableName string) *data {
	d.dataDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d data) As(alias string) *data {
	d.dataDo.DO = *(d.dataDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *data) updateTableName(table string) *data {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewUint(table, "id")
	d.UserName = field.NewString(table, "username")
	d.Password = field.NewString(table, "password")

	d.fillFieldMap()

	return d
}

func (d *data) WithContext(ctx context.Context) *dataDo { return d.dataDo.WithContext(ctx) }

func (d data) TableName() string { return d.dataDo.TableName() }

func (d data) Alias() string { return d.dataDo.Alias() }

func (d *data) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *data) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 3)
	d.fieldMap["id"] = d.ID
	d.fieldMap["username"] = d.UserName
	d.fieldMap["password"] = d.Password
}

func (d data) clone(db *gorm.DB) data {
	d.dataDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d data) replaceDB(db *gorm.DB) data {
	d.dataDo.ReplaceDB(db)
	return d
}

type dataDo struct{ gen.DO }

// SELECT * FROM @@table WHERE UserName = @username{{if username !=""}} AND password = @password{{end}}
func (d dataDo) FilterWithUsernameAndPassword(username string, password string) (result []model.Data, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, username)
	generateSQL.WriteString("SELECT * FROM data WHERE UserName = ? ")
	if username != "" {
		params = append(params, password)
		generateSQL.WriteString("AND password = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (d dataDo) Debug() *dataDo {
	return d.withDO(d.DO.Debug())
}

func (d dataDo) WithContext(ctx context.Context) *dataDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dataDo) ReadDB() *dataDo {
	return d.Clauses(dbresolver.Read)
}

func (d dataDo) WriteDB() *dataDo {
	return d.Clauses(dbresolver.Write)
}

func (d dataDo) Session(config *gorm.Session) *dataDo {
	return d.withDO(d.DO.Session(config))
}

func (d dataDo) Clauses(conds ...clause.Expression) *dataDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dataDo) Returning(value interface{}, columns ...string) *dataDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dataDo) Not(conds ...gen.Condition) *dataDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dataDo) Or(conds ...gen.Condition) *dataDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dataDo) Select(conds ...field.Expr) *dataDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dataDo) Where(conds ...gen.Condition) *dataDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dataDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *dataDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dataDo) Order(conds ...field.Expr) *dataDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dataDo) Distinct(cols ...field.Expr) *dataDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dataDo) Omit(cols ...field.Expr) *dataDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dataDo) Join(table schema.Tabler, on ...field.Expr) *dataDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dataDo) LeftJoin(table schema.Tabler, on ...field.Expr) *dataDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dataDo) RightJoin(table schema.Tabler, on ...field.Expr) *dataDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dataDo) Group(cols ...field.Expr) *dataDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dataDo) Having(conds ...gen.Condition) *dataDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dataDo) Limit(limit int) *dataDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dataDo) Offset(offset int) *dataDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *dataDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dataDo) Unscoped() *dataDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dataDo) Create(values ...*model.Data) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dataDo) CreateInBatches(values []*model.Data, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dataDo) Save(values ...*model.Data) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dataDo) First() (*model.Data, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Data), nil
	}
}

func (d dataDo) Take() (*model.Data, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Data), nil
	}
}

func (d dataDo) Last() (*model.Data, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Data), nil
	}
}

func (d dataDo) Find() ([]*model.Data, error) {
	result, err := d.DO.Find()
	return result.([]*model.Data), err
}

func (d dataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Data, err error) {
	buf := make([]*model.Data, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dataDo) FindInBatches(result *[]*model.Data, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dataDo) Attrs(attrs ...field.AssignExpr) *dataDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dataDo) Assign(attrs ...field.AssignExpr) *dataDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dataDo) Joins(fields ...field.RelationField) *dataDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dataDo) Preload(fields ...field.RelationField) *dataDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dataDo) FirstOrInit() (*model.Data, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Data), nil
	}
}

func (d dataDo) FirstOrCreate() (*model.Data, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Data), nil
	}
}

func (d dataDo) FindByPage(offset int, limit int) (result []*model.Data, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dataDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dataDo) Delete(models ...*model.Data) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dataDo) withDO(do gen.Dao) *dataDo {
	d.DO = *do.(*gen.DO)
	return d
}
