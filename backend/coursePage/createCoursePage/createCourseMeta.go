package main

import (
	"encoding/json"
	"time"

	"coursePage/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) createCourseMeta(l logger.Logger, reqBytes []byte) (course models.CourseMeta, err error) {
	err = json.Unmarshal(reqBytes, &course)
	if err != nil {
		l.Error("Course Unmarshal", err)
		return models.CourseMeta{}, err
	}

	course.Id = mongoApi.NewObjectId()
	course.ShortCard.CourseId = course.Id
	course.FullCard.CourseId = course.Id
	course.CreatedAt = time.Now()

	err = mongoApi.Insert(s.collection, course)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
