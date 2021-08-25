package models

import (
	// "fmt"
	// "net/url"
	// "strconv"
	"time"

	"github.com/jinzhu/gorm"

	// "app/db"
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

func (m LogRepository )Create() ([]Log, bool) {
	// db := db.GetDB()
	var log []Log

	return log, true
}