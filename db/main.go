package db

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	conn, err := gorm.Open("mysql", "user:user@tcp(db:3306)/log-api?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("データベース接続に失敗しました")
	}
	DB = conn
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
