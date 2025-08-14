package main

import (
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	if req.GetMethod() == suckhttp.GET {
		l.Info("Request For", req.Uri.Path[1:])

		imgBytes, err := s.getResizedImage(l, req)
		if err != nil {
			l.Error("Get image error", err)
			response = suckhttp.NewResponse(500, "Internal Server Error")
			return response, err
		}
		response = suckhttp.NewResponse(200, "Ok").SetBody(imgBytes)
	} else {
		response = suckhttp.NewResponse(405, "Method Now Allowed")
	}
	return
}

func (s *service) Close(l logger.Logger) error {
	return nil
}
