package models

import (
	// "fmt"
	// "net/url"
	// "strconv"
	"time"
	// "log"
	_ "github.com/jinzhu/gorm"

	"db"
)
type Log struct {
  ID                 int    `gorm:"AUTO_INCREMENT;primary_key"`
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

func (m LogRepository ) Create(userId int, requestData map[string]string) (bool) {
	mysql := db.GetDB()
	mysql.Create(&Log{
		UserId: userId,
		Title: requestData["Title"],
		ImageURL: requestData["ImageURL"],
		SiteURL: requestData["SiteURL"],
		OperationDatetime: time.Now(), 
	})
	return true
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

func (c LogRepository) Delete(logId int) bool {
	mysql := db.GetDB()
	log := Log{ID: logId}
	mysql.Delete(log)

	return true
}