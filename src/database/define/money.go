package main

import (
	"gorm.io/gen"
	"gorm.io/gen/helper"
)

// MoneyQuerier SQL
type MoneyQuerier interface {
	// SELECT * FROM @@table WHERE UID = @uid{{if uid != 0}}{{end}}
	FilterWithUID(uid uint32) ([]gen.T, error)
}

func GetMoney() helper.Object {
	return &Model{
		structName: "Money",
		tableName:  "Money",
		fileName:   "money",
		fields: []helper.Field{
			&ModelField{
				name:    "ID",
				typ:     "uint",
				gormTag: "column:id;type:bigint unsigned;primaryKey;autoIncrement:true",
				jsonTag: "id",
				comment: "primary key",
			},
			&ModelField{
				name:    "UID",
				typ:     "uint",
				jsonTag: "uid",
				comment: "user id",
			},
			&ModelField{
				name:    "ItemID",
				typ:     "string",
				jsonTag: "money_id",
				comment: "item id",
			},
			&ModelField{
				name:    "Amount",
				typ:     "string",
				jsonTag: "amount",
				comment: "amount",
			},
			&ModelField{
				name:    "Icon",
				typ:     "string",
				jsonTag: "icon",
				comment: "icon url(or build-in icon id)",
			},
			&ModelField{
				name:    "MoneyType",
				typ:     "string",
				jsonTag: "money_type",
				comment: "money type",
			},
		},
	}
}

func generateMoney(g *gen.Generator) {
	obj := g.GenerateModelFrom(GetMoney())
	g.ApplyBasic(obj)
	g.ApplyInterface(func(MoneyQuerier) {}, obj)
}
