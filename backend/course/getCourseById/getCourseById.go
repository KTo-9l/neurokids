package main

import (
	"course/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getCourseById(l logger.Logger, id string) (course models.Course, err error) {
	course, err = mongoApi.FindById[models.Course](s.collection, id)
	if err != nil {
		l.Error("GetById", err)
	}
	return
}
