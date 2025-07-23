package main

import (
	"encoding/json"

	"courseUser/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) createCourseUser(l logger.Logger, reqBytes []byte) (courseUser models.CourseUser, err error) {
	err = json.Unmarshal(reqBytes, &courseUser)
	if err != nil {
		l.Error("CourseUser Unmarshal", err)
		return models.CourseUser{}, err
	}

	courseUser.Id = mongoApi.NewObjectId()

	err = mongoApi.Insert(s.collection, courseUser)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
