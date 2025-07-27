package main

import (
	"coursePage/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getAllShortCourseCards(l logger.Logger) (courses []models.CourseMeta, err error) {
	selector := make(map[string]interface{})
	selector["fullCard"] = 0

	courses, err = mongoApi.ListWithSelect[models.CourseMeta](s.collection, selector)

	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
