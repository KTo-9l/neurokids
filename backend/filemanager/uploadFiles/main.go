package main

import (
	"context"

	"github.com/big-larry/mgo"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type config struct {
	DBName       string
	GridFSPrefix string
}

type service struct {
	session *mgo.Session
	bucket  *mgo.GridFS
}

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	s := &service{}

	var err error

	if err = s.initSession(c.DBName, c.GridFSPrefix); err != nil {
		return nil, nil, err
	}

	if err = s.ensurePathIndex(); err != nil {
		return nil, nil, err
	}

	return universalservice_nonepoll.CreateHTTPHandleFunc(s), s, nil
}

func main() {
	universalservice_nonepoll.InitNewServiceWithoutName(&config{}, 1)
}
