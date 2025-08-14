package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/big-larry/mgo"
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	if req.GetMethod() == suckhttp.GET {
		pathString := req.Uri.Query().Get("path")
		if pathString == "" {
			response = suckhttp.NewResponse(400, "Bad Request")
			return
		}
		path := strings.Split(pathString, "/")
		l.Debug("Path slice:", fmt.Sprint(path))

		file, err := s.getFileByPath(l, path)
		if errors.Is(err, mgo.ErrNotFound) {
			response = suckhttp.NewResponse(404, "Not Found")
			return response, nil
		} else if err != nil {
			l.Error("GetFileByPath", err)
			response = suckhttp.NewResponse(500, "Internal Server Error")
			return response, err
		}

		respBytes, err := io.ReadAll(file)
		if err != nil {
			l.Error("File To Bytes", err)
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
