package db

import (
	// "fmt"

	// "os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/lib/pq"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	fmt.Println("OK")
	// ENVIRONMENT := os.Getenv("ENVIRONMENT")
	// if ENVIRONMENT == "local" {
	// conn, err := gorm.Open("postgres", "host=db port=5432 user=user dbname=go-log-api password=user sslmode=disable")
	// DB = conn
	// if err != nil {
	// 	panic("データベース接続に失敗しました")
	// }
	// }
	// fmt.Println("ENVIRONMENT", ENVIRONMENT)
	// if ENVIRONMENT == "dev" {
	url := "postgres://qhpyygbsdauxkr:cad56e5a1b0b1d118987fb8941090713803c61d911eae9b209671854bab37388@ec2-35-171-250-21.compute-1.amazonaws.com:5432/d1sih1b00rn14v"
	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=require"

	conn, err := gorm.Open("postgres", connection)
	DB = conn
	// }

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
