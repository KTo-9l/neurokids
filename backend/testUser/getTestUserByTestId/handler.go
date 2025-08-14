package main

import (
	"auth_helpers"
	"dikobra3/mongoApi"
	"dikobra3/utils"
	"errors"
	"strconv"

	"github.com/big-larry/mgo"
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	perm, err := auth_helpers.GetPerms[auth_helpers.Perms](req)
	if err != nil {
		response = suckhttp.NewResponse(403, "forbidden")
		return
	}
	if p, ok := perm.Perms[0]; !ok || p&auth_helpers.AllPerms == 0 {
		response = suckhttp.NewResponse(403, "forbidden")
		return
	}

	if req.GetMethod() == suckhttp.GET {
		uidStr := req.GetHeader("x-user-id")
		var uid int
		if uid, err = strconv.Atoi(uidStr); err != nil {
			response = suckhttp.NewResponse(401, "Unauthorized")
			return
		}

		testId := req.Uri.Query().Get("testId")
		if !mongoApi.IsObjectId(testId) {
			response = suckhttp.NewResponse(400, "Bad Request")
			return
		}

		testUser, err := s.getTestUserByTestId(l, uid, testId)
		if errors.Is(err, mgo.ErrNotFound) {
			response = suckhttp.NewResponse(404, "Not Found")
			return response, nil
		} else if err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
		} else {
			body, err := utils.ObjectToBytes(testUser)
			if err != nil {
				l.Error("ObjectToBytes", err)
				response = suckhttp.NewResponse(500, "Internal Server Error")
				return response, err
			}
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
