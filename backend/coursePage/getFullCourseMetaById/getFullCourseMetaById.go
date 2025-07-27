package main

import (
	"coursePage/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getFullCourseMetaById(l logger.Logger, id string) (courseMeta models.CourseMeta, err error) {
	courseMeta, err = mongoApi.FindById[models.CourseMeta](s.collection, id)
	if err != nil {
		l.Error("GetById", err)
	}
	return
}
