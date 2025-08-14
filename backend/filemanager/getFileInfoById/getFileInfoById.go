package main

import (
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getFileInfoById(l logger.Logger, id string) (mongoApi.GridFSFile, error) {
	file, err := mongoApi.FindById[mongoApi.GridFSFile](s.bucket.Files, id)
	if err != nil {
		l.Error("get query error", err)
		return mongoApi.GridFSFile{}, err
	}
	return file, nil
}
