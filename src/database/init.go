package database

import (
	"MoneyManager/src/model"
	"MoneyManager/src/query"
	"MoneyManager/src/utility/log"
	"MoneyManager/src/utility/utils"
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var initFunc map[string]func(*model.Config) *gorm.DB = map[string]func(*model.Config) *gorm.DB{
	"mysql":  initMysql,
	"sqlite": initSqlite,
}

func initMysql(config *model.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database.UserName, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.TrackError(context.TODO(), err)
		panic(err)
	}

	return db
}

func initSqlite(config *model.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.Database.FilePath), &gorm.Config{})
	if err != nil {
		log.TrackError(context.TODO(), err)
		panic(err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

func GetQuery() *query.Query {
	return query.Use(db)
}

func Init() {
	config := utils.GetConfig()

	if f, ok := initFunc[config.Database.Type]; ok {
		db = f(config)
	} else {
		panic("Database type not found")
	}

}
