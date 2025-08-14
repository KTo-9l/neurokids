package main

import (
	"auth_helpers"
	"strconv"

	"dikobra3/mongoApi"
	"dikobra3/utils"

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
		chatId := req.Uri.Query().Get("chatId")
		if chatId == "" || !mongoApi.IsObjectId(chatId) {
			response = suckhttp.NewResponse(400, "Bad Request")
			return response, err
		}

		amountString := req.Uri.Query().Get("amount")
		amount := 0
		if amountString != "" {
			amount, err = strconv.Atoi(amountString)
			if err != nil {
				response = suckhttp.NewResponse(400, "Bad Request")
				return response, err
			}
		}

		if messages, err := s.getMessages(l, chatId, amount); err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
		} else {
			body, err := utils.ObjectToBytes(messages)
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
