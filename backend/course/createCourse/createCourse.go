package main

import (
	"encoding/json"
	"time"

	"course/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) createCourse(l logger.Logger, reqBytes []byte) (course models.Course, err error) {
	err = json.Unmarshal(reqBytes, &course)
	if err != nil {
		l.Error("Course Unmarshal", err)
		return models.Course{}, err
	}

	course.Id = mongoApi.NewObjectId()
	course.CreatedAt = time.Now()

	err = mongoApi.Insert(s.collection, course)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
