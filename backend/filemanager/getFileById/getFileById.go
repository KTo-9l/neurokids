package main

import (
	"dikobra3/mongoApi"

	"github.com/big-larry/mgo"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getFileById(l logger.Logger, id string) (*mgo.GridFile, error) {
	file, err := mongoApi.GetFileById(s.bucket, id)
	if err != nil {
		l.Error("get query error", err)
		return nil, err
	}
	return file, nil
}
