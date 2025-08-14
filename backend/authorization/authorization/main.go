package main

import (
	"authorization/database"
	"context"

	"github.com/okonma-violet/services/logs/logger"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type config struct {
	ConnectionString string
}

type service struct {
	rep *database.Repo
}

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	rep, err := database.OpenRepository(c.ConnectionString)
	if err != nil {
		return nil, nil, err
	}
	s := &service{
		rep: rep,
	}
	return universalservice_nonepoll.CreateHTTPHandleFunc(s), s, nil
}

func (s *service) Close(l logger.Logger) error {
	s.rep.Close()
	return nil
}

func main() {
	universalservice_nonepoll.InitNewServiceWithoutName(&config{}, 1)
}
