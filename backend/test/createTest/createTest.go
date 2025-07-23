package main

import (
	"encoding/json"

	"test/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) createTest(l logger.Logger, reqBytes []byte) (test models.Test, err error) {
	err = json.Unmarshal(reqBytes, &test)
	if err != nil {
		l.Error("Test Unmarshal", err)
		return models.Test{}, err
	}

	test.Id = mongoApi.NewObjectId()

	err = mongoApi.Insert(s.collection, test)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
