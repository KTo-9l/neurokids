package main

import (
	"dikobra3/mongoApi"

	"github.com/big-larry/mgo"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getFileByPath(l logger.Logger, path []string) (*mgo.GridFile, error) {
	file, err := mongoApi.GetFileByPath(s.bucket, path)
	if err != nil {
		l.Error("get query error", err)
		return nil, err
	}
	return file, nil
}
