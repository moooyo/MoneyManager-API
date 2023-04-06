package main

import (
	"gorm.io/gen"
	"gorm.io/gen/helper"
)

// Dynamic SQL
type UserQuerier interface {
	// SELECT * FROM @@table WHERE UserName = @username{{if username !=""}} AND password = @password{{end}}
	FilterWithUsernameAndPassword(username, password string) ([]gen.T, error)
}

func GetUser() helper.Object {
	return &Model{
		structName: "Data",
		tableName:  "data",
		fileName:   "user",
		fields: []helper.Field{
			&ModelField{
				name:    "ID",
				typ:     "uint",
				gormTag: "column:id;type:bigint unsigned;primaryKey;autoIncrement:true",
				jsonTag: "id",
				comment: "key id",
			},
			&ModelField{
				name:    "UserName",
				typ:     "string",
				jsonTag: "user_name",
				comment: "username",
			},
			&ModelField{
				name:    "Password",
				typ:     "string",
				jsonTag: "password",
				comment: "password",
			},
		},
	}
}

func generateUser(g *gen.Generator) {
	obj := g.GenerateModelFrom(GetUser())
	g.ApplyBasic(obj)
	g.ApplyInterface(func(UserQuerier) {}, obj)
}
