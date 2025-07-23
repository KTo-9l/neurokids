package main

import (
	"encoding/json"

	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) initSession(dbName, prefix string) (err error) {
	s.session, err = mongoApi.Connect()
	if err != nil {
		return err
	}
	s.bucket = s.session.DB(dbName).GridFS(prefix)
	return
}

func (s *service) ensurePathIndex() error {
	return mongoApi.EnsurePathIndex(s.bucket)
}

func structToBytes(l logger.Logger, str interface{}) ([]byte, error) {
	bytes, err := json.Marshal(str)
	if err != nil {
		l.Error("[error marshal files]", err)
		return nil, err
	}
	return bytes, nil
}
