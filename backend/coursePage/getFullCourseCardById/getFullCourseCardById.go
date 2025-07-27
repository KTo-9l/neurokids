package main

import (
	"coursePage/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getFullCourseCardById(l logger.Logger, id string) (courseMeta models.CourseMeta, err error) {
	selector := make(map[string]interface{})
	selector["shortCard"] = 0

	courseMeta, err = mongoApi.FindByIdWithSelect[models.CourseMeta](s.collection, id, selector)
	if err != nil {
		l.Error("GetById", err)
	}
	return
}
