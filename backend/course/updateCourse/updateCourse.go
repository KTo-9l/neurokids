package main

import (
	"encoding/json"

	"course/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateCourse(l logger.Logger, reqBytes []byte) (course models.Course, err error) {
	err = json.Unmarshal(reqBytes, &course)
	if err != nil {
		l.Error("Course Unmarshal", err)
		return models.Course{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"title":   course.Title,
			"lessons": course.Lessons,
			"deleted": course.Deleted,
		},
	}

	err = mongoApi.UpdateById(s.collection, course.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
