package main

import (
	"courseUser/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getCourseUser(l logger.Logger, uid string) (coursesUser []models.CourseUser, err error) {
	selector := make(map[string]interface{})
	selector["userId"] = mongoApi.StringToObjectId(uid)

	coursesUser, err = mongoApi.ListTypifiedCollectionWithSelector[models.CourseUser](s.collection, selector)
	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
