package main

import (
	"chat/models"
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getChats(l logger.Logger, uid string) (chats []models.Chat, err error) {
	selector := make(map[string]interface{})
	selector["members"] = mongoApi.StringToObjectId(uid)

	chats, err = mongoApi.ListTypifiedCollectionWithSelector[models.Chat](s.collection, selector)
	if err != nil {
		l.Error("ListAllCollection", err)
	}
	return
}
