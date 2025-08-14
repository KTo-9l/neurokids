package main

import (
	"context"
	"dikobra3/mongoApi"

	"github.com/big-larry/mgo"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type config struct {
	ConnectionString string
}

type service struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	s := &service{}

	var err error

	if s.session, err = mongoApi.Connect(c.ConnectionString); err != nil {
		return nil, nil, err
	}
	s.collection = s.session.DB("test").C("testUser")

	if err = mongoApi.EnsureIndexKey(s.collection, "userId"); err != nil {
		return nil, nil, err
	}

	return universalservice_nonepoll.CreateHTTPHandleFunc(s), s, nil
}

func main() {
	universalservice_nonepoll.InitNewServiceWithoutName(&config{}, 1)
}
