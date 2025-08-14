package main

import (
	"dikobra3/utils"
	"errors"

	"github.com/big-larry/mgo"
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	if req.GetMethod() == suckhttp.GET {
		id := req.Uri.Query().Get("id")
		if id == "" {
			response = suckhttp.NewResponse(400, "Bad Request")
			return
		}

		file, err := s.getFileInfoById(l, id)
		if errors.Is(err, mgo.ErrNotFound) {
			response = suckhttp.NewResponse(404, "Not Found")
			return response, nil
		} else if err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
			return response, err
		}

		body, err := utils.ObjectToBytes(file)
		if err != nil {
			l.Error("Object To Bytes", err)
			response = suckhttp.NewResponse(500, "Internal Server Error")
			return response, err
		} else {
			response = suckhttp.NewResponse(200, "OK").AddHeader("Content-Type", "application/json").SetBody(body)
		}
	} else {
		response = suckhttp.NewResponse(405, "Method Not Allowed")
	}
	return
}

func (s *service) Close(l logger.Logger) error {
	s.session.Close()
	return nil
}
