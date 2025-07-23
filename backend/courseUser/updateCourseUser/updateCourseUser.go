package main

import (
	"encoding/json"

	"courseUser/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateCourseUser(l logger.Logger, reqBytes []byte) (courseUser models.CourseUser, err error) {
	err = json.Unmarshal(reqBytes, &courseUser)
	if err != nil {
		l.Error("CourseUser Unmarshal", err)
		return models.CourseUser{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"purchased": courseUser.Purchased,
			"progress": bson.M{
				"opened":   courseUser.Progress.Opened,
				"stage":    courseUser.Progress.Stage,
				"finished": courseUser.Progress.Finished,
			},
		},
	}

	err = mongoApi.UpdateById(s.collection, courseUser.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
