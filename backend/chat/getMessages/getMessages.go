package main

import (
	"chat/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getMessages(l logger.Logger, chatId string, amount int) (messages []models.Message, err error) {
	selector := make(map[string]interface{})
	selector["chatId"] = mongoApi.StringToObjectId(chatId)

	if amount == 0 {
		messages, err = mongoApi.ListTypifiedSortedCollectionByFields[models.Message](s.collection, selector, []string{"-time"})
	} else {
		messages, err = mongoApi.ListLimitedCollectionByFields[models.Message](s.collection, selector, []string{"-time"}, amount)
	}
	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
