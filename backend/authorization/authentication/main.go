package main

import (
	"authorization/database"
	"context"
	"errors"
	"text/template"

	"github.com/okonma-violet/services/logs/logger"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

// read this from configfile
type config struct {
	HtmlIndexTemplate string
	LayoutTemplate    string
	ConnectionString  string
}

// your shit here
type service struct {
	rep            *database.Repo
	index_template *template.Template
	// pub_storage     *universalservice_nonepoll.Publisher
}

// const subServiceName universalservice_nonepoll.ServiceName = "journal_insert"

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	if c.HtmlIndexTemplate == "" || c.LayoutTemplate == "" {
		return nil, nil, errors.New("no fields \"HtmlTemplate\" or \"LayoutTemplate\" in a config file")
	}
	t, err := template.New("index.html").Funcs(template.FuncMap{"escape": template.HTMLEscapeString}).ParseFiles(c.HtmlIndexTemplate, c.LayoutTemplate)
	if err != nil {
		return nil, nil, err
	}
	rep, err := database.OpenRepository(c.ConnectionString)
	if err != nil {
		return nil, nil, err
	}
	s := &service{
		index_template: t,
		rep:            rep,
		// pub_storage:     pubs_getter.Get(subServiceName),
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
