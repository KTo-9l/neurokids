package main

import (
	"testUser/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getAllTestsUser(l logger.Logger, uid string) (testsUser []models.TestUser, err error) {
	selector := make(map[string]interface{})
	selector["userId"] = mongoApi.StringToObjectId(uid)

	testsUser, err = mongoApi.ListTypifiedCollectionWithSelector[models.TestUser](s.collection, selector)
	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
