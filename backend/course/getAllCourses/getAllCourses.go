package main

import (
	"course/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getAllCourses(l logger.Logger) (courses []models.Course, err error) {
	courses, err = mongoApi.ListTypifiedCollection[models.Course](s.collection)
	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
