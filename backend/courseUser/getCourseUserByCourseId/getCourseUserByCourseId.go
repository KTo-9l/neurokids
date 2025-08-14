package main

import (
	"courseUser/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getCourseUserByCourseId(l logger.Logger, uid int, courseId string) (coursesUser models.CourseUser, err error) {
	selector := make(map[string]interface{})
	selector["userId"] = uid
	selector["courseId"] = mongoApi.StringToObjectId(courseId)

	coursesUser, err = mongoApi.FindOne[models.CourseUser](s.collection, selector)
	if err != nil {
		l.Error("ListCollection", err)
	}
	return
}
