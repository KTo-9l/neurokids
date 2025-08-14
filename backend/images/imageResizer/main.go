package main

import (
	"context"

	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type config struct {
}

type service struct {
	pubs_getter universalservice_nonepoll.Publishers_getter
}

const (
	filesharing_get_file_by_path = "filesharing_get_file_by_path"
)

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	s := &service{
		pubs_getter: pubs_getter,
	}

	return universalservice_nonepoll.CreateHTTPHandleFunc(s), s, nil
}

func main() {
	services := []universalservice_nonepoll.ServiceName{
		filesharing_get_file_by_path,
	}

	universalservice_nonepoll.InitNewServiceWithoutName(&config{}, 1, services...)
}
