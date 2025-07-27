package main

import (
	"encoding/json"

	"coursePage/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateCourseMeta(l logger.Logger, reqBytes []byte) (course models.CourseMeta, err error) {
	err = json.Unmarshal(reqBytes, &course)
	if err != nil {
		l.Error("Course Unmarshal", err)
		return models.CourseMeta{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"title":   course.Title,
			"author":  course.Author,
			"shown":   course.Shown,
			"deleted": course.Deleted,
		},
	}

	err = mongoApi.UpdateById(s.collection, course.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
