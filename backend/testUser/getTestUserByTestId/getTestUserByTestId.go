package main

import (
	"courseUser/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getTestUserByTestId(l logger.Logger, uid int, testId string) (testUser models.TestUser, err error) {
	selector := make(map[string]interface{})
	selector["userId"] = uid
	selector["testId"] = mongoApi.StringToObjectId(testId)

	testUser, err = mongoApi.FindOne[models.TestUser](s.collection, selector)
	if err != nil {
		l.Error("ListCollection", err)
	}
	return
}
