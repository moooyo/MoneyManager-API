package main

import (
	"gorm.io/gen"
	"gorm.io/gen/helper"
)

// Dynamic SQL
type CommentQuerier interface {
	// SELECT * FROM @@table WHERE ID = @id{{if id !=""}}{{end}}
	FilterWithID(id string) ([]gen.T, error)
}

func GetComment() helper.Object {
	return &Model{
		structName: "Comment",
		tableName:  "comment",
		fileName:   "comment",
		fields: []helper.Field{
			&ModelField{
				name:    "ID",
				typ:     "uint",
				gormTag: "column:id;type:bigint unsigned;primaryKey;autoIncrement:true",
				jsonTag: "id",
				comment: "主键",
			},
			&ModelField{
				name:    "Comment",
				typ:     "string",
				jsonTag: "user_name",
				comment: "账户名",
			},
		},
	}
}

func generateComment(g *gen.Generator) {
	obj := g.GenerateModelFrom(GetComment())
	g.ApplyBasic(obj)
	g.ApplyInterface(func(CommentQuerier) {}, obj)
}
