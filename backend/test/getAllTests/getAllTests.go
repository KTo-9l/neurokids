package main

import (
	"test/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getAllTests(l logger.Logger) (tests []models.Test, err error) {
	tests, err = mongoApi.ListTypifiedCollection[models.Test](s.collection)
	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
