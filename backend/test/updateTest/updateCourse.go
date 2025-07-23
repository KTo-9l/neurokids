package main

import (
	"encoding/json"

	"test/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateTest(l logger.Logger, reqBytes []byte) (test models.Test, err error) {
	err = json.Unmarshal(reqBytes, &test)
	if err != nil {
		l.Error("Test Unmarshal", err)
		return models.Test{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"title":     test.Title,
			"questions": test.Questions,
			"deleted":   test.Deleted,
		},
	}

	err = mongoApi.UpdateById(s.collection, test.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
