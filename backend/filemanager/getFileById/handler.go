package main

import (
	"fmt"

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

		file, err := s.getFileById(l, id)
		if err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
			return response, err
		}

		respBytes, err := fileToBytes(l, file)
		if err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
		} else {
			response = suckhttp.NewResponse(200, "OK").
				AddHeader("Content-Type", "application/octet-stream").
				AddHeader("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.Name())).
				SetBody(respBytes)
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
