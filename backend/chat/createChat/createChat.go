package main

import (
	"encoding/json"

	"chat/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) createChat(l logger.Logger, reqBytes []byte) (chat models.Chat, err error) {
	err = json.Unmarshal(reqBytes, &chat)
	if err != nil {
		l.Error("Chat Unmarshal", err)
		return models.Chat{}, err
	}

	chat.Id = mongoApi.NewObjectId()

	err = mongoApi.Insert(s.collection, chat)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
