package main

import (
	"dikobra3/mongoApi"
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
