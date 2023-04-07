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

func newMoney(db *gorm.DB, opts ...gen.DOOption) money {
	_money := money{}

	_money.moneyDo.UseDB(db, opts...)
	_money.moneyDo.UseModel(&model.Money{})

	tableName := _money.moneyDo.TableName()
	_money.ALL = field.NewAsterisk(tableName)
	_money.ID = field.NewUint(tableName, "id")
	_money.UID = field.NewUint(tableName, "uid")
	_money.ItemID = field.NewString(tableName, "itemid")
	_money.Amount = field.NewString(tableName, "amount")
	_money.Icon = field.NewString(tableName, "icon")
	_money.MoneyType = field.NewString(tableName, "moneytype")

	_money.fillFieldMap()

	return _money
}

type money struct {
	moneyDo moneyDo

	ALL       field.Asterisk
	ID        field.Uint   // primary key
	UID       field.Uint   // user id
	ItemID    field.String // item id
	Amount    field.String // amount
	Icon      field.String // icon url(or build-in icon id)
	MoneyType field.String // money type

	fieldMap map[string]field.Expr
}

func (m money) Table(newTableName string) *money {
	m.moneyDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m money) As(alias string) *money {
	m.moneyDo.DO = *(m.moneyDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *money) updateTableName(table string) *money {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.UID = field.NewUint(table, "uid")
	m.ItemID = field.NewString(table, "itemid")
	m.Amount = field.NewString(table, "amount")
	m.Icon = field.NewString(table, "icon")
	m.MoneyType = field.NewString(table, "moneytype")

	m.fillFieldMap()

	return m
}

func (m *money) WithContext(ctx context.Context) *moneyDo { return m.moneyDo.WithContext(ctx) }

func (m money) TableName() string { return m.moneyDo.TableName() }

func (m money) Alias() string { return m.moneyDo.Alias() }

func (m *money) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *money) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["uid"] = m.UID
	m.fieldMap["itemid"] = m.ItemID
	m.fieldMap["amount"] = m.Amount
	m.fieldMap["icon"] = m.Icon
	m.fieldMap["moneytype"] = m.MoneyType
}

func (m money) clone(db *gorm.DB) money {
	m.moneyDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m money) replaceDB(db *gorm.DB) money {
	m.moneyDo.ReplaceDB(db)
	return m
}

type moneyDo struct{ gen.DO }

// SELECT * FROM @@table WHERE UID = @uid{{if uid != 0}}{{end}}
func (m moneyDo) FilterWithUID(uid uint32) (result []model.Money, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	generateSQL.WriteString("SELECT * FROM Money WHERE UID = ? ")
	if uid != 0 {
	}

	var executeSQL *gorm.DB
	executeSQL = m.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (m moneyDo) Debug() *moneyDo {
	return m.withDO(m.DO.Debug())
}

func (m moneyDo) WithContext(ctx context.Context) *moneyDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m moneyDo) ReadDB() *moneyDo {
	return m.Clauses(dbresolver.Read)
}

func (m moneyDo) WriteDB() *moneyDo {
	return m.Clauses(dbresolver.Write)
}

func (m moneyDo) Session(config *gorm.Session) *moneyDo {
	return m.withDO(m.DO.Session(config))
}

func (m moneyDo) Clauses(conds ...clause.Expression) *moneyDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m moneyDo) Returning(value interface{}, columns ...string) *moneyDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m moneyDo) Not(conds ...gen.Condition) *moneyDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m moneyDo) Or(conds ...gen.Condition) *moneyDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m moneyDo) Select(conds ...field.Expr) *moneyDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m moneyDo) Where(conds ...gen.Condition) *moneyDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m moneyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *moneyDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m moneyDo) Order(conds ...field.Expr) *moneyDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m moneyDo) Distinct(cols ...field.Expr) *moneyDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m moneyDo) Omit(cols ...field.Expr) *moneyDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m moneyDo) Join(table schema.Tabler, on ...field.Expr) *moneyDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m moneyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *moneyDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m moneyDo) RightJoin(table schema.Tabler, on ...field.Expr) *moneyDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m moneyDo) Group(cols ...field.Expr) *moneyDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m moneyDo) Having(conds ...gen.Condition) *moneyDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m moneyDo) Limit(limit int) *moneyDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m moneyDo) Offset(offset int) *moneyDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m moneyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *moneyDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m moneyDo) Unscoped() *moneyDo {
	return m.withDO(m.DO.Unscoped())
}

func (m moneyDo) Create(values ...*model.Money) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m moneyDo) CreateInBatches(values []*model.Money, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m moneyDo) Save(values ...*model.Money) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m moneyDo) First() (*model.Money, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Money), nil
	}
}

func (m moneyDo) Take() (*model.Money, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Money), nil
	}
}

func (m moneyDo) Last() (*model.Money, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Money), nil
	}
}

func (m moneyDo) Find() ([]*model.Money, error) {
	result, err := m.DO.Find()
	return result.([]*model.Money), err
}

func (m moneyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Money, err error) {
	buf := make([]*model.Money, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m moneyDo) FindInBatches(result *[]*model.Money, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m moneyDo) Attrs(attrs ...field.AssignExpr) *moneyDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m moneyDo) Assign(attrs ...field.AssignExpr) *moneyDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m moneyDo) Joins(fields ...field.RelationField) *moneyDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m moneyDo) Preload(fields ...field.RelationField) *moneyDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m moneyDo) FirstOrInit() (*model.Money, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Money), nil
	}
}

func (m moneyDo) FirstOrCreate() (*model.Money, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Money), nil
	}
}

func (m moneyDo) FindByPage(offset int, limit int) (result []*model.Money, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m moneyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m moneyDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m moneyDo) Delete(models ...*model.Money) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *moneyDo) withDO(do gen.Dao) *moneyDo {
	m.DO = *do.(*gen.DO)
	return m
}
