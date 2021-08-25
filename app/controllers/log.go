package controllers

import (
	"app/models"
)

type Log struct{}

func NewLog() Log {
	return Log{}
}

func (c Log) CreateLog(userId int) interface{} {
	repo := models.NewLogRepository()
	log, result := repo.Create(userId)
	if result {
		return log
	}
	return result
}

func (c Log) GetLogList(userId int) interface{} {
	repo := models.NewLogRepository()
	logs, err := repo.GetAll(userId)
	if err == false {
		panic("err")
	}

	return logs
}