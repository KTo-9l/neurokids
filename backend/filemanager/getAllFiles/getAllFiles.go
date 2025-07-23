package main

import (
	"dikobra3/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getAllFiles(l logger.Logger) ([]mongoApi.GridFSFile, error) {
	files, err := mongoApi.ListAllFiles(s.bucket)
	if err != nil {
		l.Error("[get query error]", err)
		return nil, err
	}
	return files, nil
}
