package main

import (
	"test/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getTestById(l logger.Logger, id string) (test models.Test, err error) {
	test, err = mongoApi.FindById[models.Test](s.collection, id)
	if err != nil {
		l.Error("GetById", err)
	}
	return
}
