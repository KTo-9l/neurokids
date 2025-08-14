package main

import (
	"authorization/auth_helpers"
	"net/url"
	"strconv"

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

	if req.GetMethod() == suckhttp.POST {
		uidStr := req.GetHeader("x-user-id")
		var uid int
		if uid, err = strconv.Atoi(uidStr); err != nil {
			response = suckhttp.NewResponse(401, "Unauthorized")
			return
		}

		var form url.Values
		form, err = auth_helpers.ParseForm(req)
		if err != nil {
			response = suckhttp.NewResponse(500, "Server error")
			return
		}
		var name, email = "", ""
		name, _ = auth_helpers.TryGetFormRawValue(form, "name")
		email, _ = auth_helpers.TryGetFormRawValue(form, "email")

		if name == "" || email == "" {
			l.Debug("Bad formdata", "Can't get name or email")
			response = suckhttp.NewResponse(404, "Bad Request")
			return
		}

		if err := s.updateInfo(l, uid, name, email); err != nil {
			l.Error("Error update password", err)
			response = suckhttp.NewResponse(500, "Internal Server Error")
		} else {
			response = suckhttp.NewResponse(200, "OK")
		}
	} else {
		response = suckhttp.NewResponse(405, "Method Not Allowed")
	}
	return
}
