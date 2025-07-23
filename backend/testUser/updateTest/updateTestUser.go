package main

import (
	"encoding/json"

	"testUser/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateTestUser(l logger.Logger, reqBytes []byte) (testUser models.TestUser, err error) {
	err = json.Unmarshal(reqBytes, &testUser)
	if err != nil {
		l.Error("TestUser Unmarshal", err)
		return models.TestUser{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"progress": bson.M{
				"opened":   testUser.Progress.Opened,
				"stage":    testUser.Progress.Stage,
				"finished": testUser.Progress.Finished,
				"correct":  testUser.Progress.Correct,
			},
		},
	}

	err = mongoApi.UpdateById(s.collection, testUser.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
