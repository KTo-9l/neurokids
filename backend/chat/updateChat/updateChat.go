package main

import (
	"encoding/json"

	"chat/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateChat(l logger.Logger, reqBytes []byte) (chat models.Chat, err error) {
	err = json.Unmarshal(reqBytes, &chat)
	if err != nil {
		l.Error("Chat Unmarshal", err)
		return models.Chat{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"title":   chat.Title,
			"members": chat.Members,
			"isGroup": chat.IsGroup,
			"deleted": chat.Deleted,
		},
	}

	err = mongoApi.UpdateById(s.collection, chat.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
