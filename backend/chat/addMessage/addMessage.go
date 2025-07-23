package main

import (
	"encoding/json"

	"chat/models"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) addMessage(l logger.Logger, reqBytes []byte) (message models.Message, err error) {
	err = json.Unmarshal(reqBytes, &message)
	if err != nil {
		l.Error("Message Unmarshal", err)
		return models.Message{}, err
	}

	message.Id = mongoApi.NewObjectId()

	err = mongoApi.Insert(s.collection, message)
	if err != nil {
		l.Error("coll.Insert", err)
	}
	return
}
