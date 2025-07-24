package main

import (
	"encoding/json"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func initHub() {
	hub = &Hub{
		clients: make(map[bson.ObjectId][]*Client),
	}
}

func structToBytes(l logger.Logger, str interface{}) ([]byte, error) {
	bytes, err := json.Marshal(str)
	if err != nil {
		l.Error("Marshal answer", err)
		return nil, err
	}
	return bytes, nil
}
