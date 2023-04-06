package main

import (
	"strings"

	"gorm.io/gen/helper"
)

var _ helper.Object = new(Model)

type Model struct {
	structName string
	tableName  string
	fileName   string
	fields     []helper.Field
}

// TableName return table name
func (d *Model) TableName() string { return d.tableName }

// StructName return struct name
func (d *Model) StructName() string { return d.structName }

// FileName return file name
func (d *Model) FileName() string { return d.fileName }

// ImportPkgPaths return import package paths
func (d *Model) ImportPkgPaths() []string { return nil }

// Fields return fields
func (d *Model) Fields() []helper.Field { return d.fields }

// DemoField demo field
type ModelField struct {
	name    string
	typ     string
	gormTag string
	jsonTag string
	tag     string
	comment string
}

// Name return name
func (f *ModelField) Name() string { return f.name }

// Type return field type
func (f *ModelField) Type() string { return f.typ }

// ColumnName return column name
func (f *ModelField) ColumnName() string { return strings.ToLower(f.name) }

// GORMTag return gorm tag
func (f *ModelField) GORMTag() string { return f.gormTag }

// JSONTag return json tag
func (f *ModelField) JSONTag() string { return f.jsonTag }

// Tag return new tag
func (f *ModelField) Tag() string { return f.tag }

// Comment return comment
func (f *ModelField) Comment() string { return f.comment }
