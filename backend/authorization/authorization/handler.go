package main

import (
	"authorization/auth_helpers"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

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

	perms, err1 := AuthorizationUser(s.rep, UserToken(result.Token), UserId(result.Uid))
	// fmt.Println(token, result.Token, perms)
	if err1 != nil {
		err = err1
		return
	}
	if len(perms) == 0 {
		response = suckhttp.NewResponse(403, "Forbidden")
		return
	}
	body, err1 := json.Marshal(auth_helpers.AuthResultWithPerms[auth_helpers.Perms]{AuthResult: *result, Perms: perms})
	if err1 != nil {
		err = err1
		return
	}
	// uri, err1 := helpers.GetOriginalUri(req)
	// if err1 != nil {
	// 	err = err1
	// 	return
	// }
	// reqid, err1 := helpers.GetRequestId(req)
	// if err1 != nil {
	// 	err = err1
	// 	return
	// }
	response = suckhttp.
		NewResponse(200, "OK").
		// NewResponse(301, "OK"). // hack for nginx
		AddHeader("x-perm", string(body)).
		AddHeader("x-user-id", strconv.Itoa(result.Uid))
		// SetBody(body).

	return
}
