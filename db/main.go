package db

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	conn, err := gorm.Open("postgres", "host=db port=5432 user=user dbname=go-log-api password=user sslmode=disable")
	if err != nil {
		panic("データベース接続に失敗しました")
	}
	DB = conn
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
