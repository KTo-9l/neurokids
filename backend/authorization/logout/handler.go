package main

import (
	"authorization/auth_helpers"
	"errors"
	"net/url"

	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	response = suckhttp.NewResponse(500, "Server error")
	var form url.Values
	form, err = auth_helpers.ParseForm(req)
	if err != nil {
		response = suckhttp.NewResponse(500, "Server error")
		return
	}

	token := form.Get("token")
	if token == "" {
		token, _ = req.GetCookie("token")
	}
	if token == "" {
		response = suckhttp.NewResponse(403, "Forbidden")
		err = errors.New("token is empty")
		return
	}
	result, err1 := auth_helpers.OpenToken(token)
	if err1 != nil {
		err = err1
		return
	}

	err = LogoutUser(s.rep, result.Uid)
	if err != nil {
		l.Error("Error logout user", err)
		return
	}

	response = suckhttp.NewResponse(302, "Redirect")
	// response = suckhttp.NewResponse(200, "Ok")
	response.AddHeader(suckhttp.Set_Cookie, "token=")
	response.AddHeader(suckhttp.Location, "/neurokids")
	return
}
