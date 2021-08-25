package models

import (
	// "fmt"
	// "net/url"
	// "strconv"
	"time"

	"github.com/jinzhu/gorm"

	"app/db"
)
type Log struct {
  gorm.Model
  UserId        	 int 
  Title 			 string
  ImageURL 			 string
  SiteURL 			 string
  OperationDatetime  time.Time
}

type LogRepository struct{}

func NewLogRepository() LogRepository {
	return LogRepository{}
}

func (m LogRepository ) Create(userId int) ([]Log, bool) {
	// db := db.GetDB()
	var log []Log

	return log, true
}

func (c LogRepository) GetAll(userId int) ([]Log, bool) {
	mysql := db.GetDB()
	var logs []Log
	// queryLogs := Log{UserId: userId}
	// state := mysql.Table("logs").Where(queryLogs).Group("id")
	mysql.Table("logs").Where("user_id = ?", userId).Scan(&logs)
	// mysql.Where(map[string]interface{}{"user_id": userId}).Find(&logs)
	// logs, err := mysql.Query("SELECT * FROM logs where user_id=?", userId)
	// if err != nil {
	// 	panic("err")
	// }
	// state.Find(&logs)

	return logs, true
}