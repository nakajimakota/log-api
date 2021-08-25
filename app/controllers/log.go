package controllers

import (
	"app/models"
	// "net/url"
)

type Log struct{}

func NewLog() Log {
	return Log{}
}

func (c Log) CreateLog() interface{} {
	repo := models.NewLogRepository()
	log, result := repo.Create()
	if result {
		return log
	}
	return result
}