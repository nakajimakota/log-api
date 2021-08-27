package controllers

import (
	"app/models"
)

type Log struct{}

func NewLog() Log {
	return Log{}
}

func (c Log) CreateLog(userId int, requestData map[string]string) interface{} {
	repo := models.NewLogRepository()
	err := repo.Create(userId, requestData)
	if err != false {
		return false
	}
	return true
}

func (c Log) GetLogList(userId int) interface{} {
	repo := models.NewLogRepository()
	logs, err := repo.GetAll(userId)
	if err == false {
		panic("err")
	}

	return logs
}

func (c Log) DeleteLog(logId int) interface{} {
	repo := models.NewLogRepository()
	res := repo.Delete(logId)
	if res == false {
		panic("err")
	}

	return true
}