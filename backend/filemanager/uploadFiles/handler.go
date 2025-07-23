package main

import (
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	if req.GetMethod() == suckhttp.POST {
		respContent, err := s.uploadFiles(l, req)
		if respContent == nil && err == nil {
			response = suckhttp.NewResponse(400, "Bad Request")
			return response, err
		}
		if err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
			return response, err
		}

		respBytes, err := structToBytes(l, respContent)
		if err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
		} else {
			response = suckhttp.NewResponse(200, "OK").AddHeader("Content-Type", "application/json").SetBody(respBytes)
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
