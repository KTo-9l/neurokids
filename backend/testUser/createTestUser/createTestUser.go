package main

import (
	"encoding/json"

	"testUser/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) createTestUser(l logger.Logger, reqBytes []byte) (testUser models.TestUser, err error) {
	err = json.Unmarshal(reqBytes, &testUser)
	if err != nil {
		l.Error("TestUser Unmarshal", err)
		return models.TestUser{}, err
	}

	testUser.Id = mongoApi.NewObjectId()

	err = mongoApi.Insert(s.collection, testUser)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
