package main

import (
	"authorization/auth_helpers"
	"errors"
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
		var oldPassword, newPassword = "", ""
		oldPassword, _ = auth_helpers.TryGetFormRawValue(form, "old_password")
		newPassword, _ = auth_helpers.TryGetFormRawValue(form, "new_password")

		if oldPassword == newPassword {
			err = errors.New("old and new password must be different")
			l.Error("Same passwords", err)
			response = suckhttp.NewResponse(404, "Bad Request")
			return
		}
		if oldPassword == "" || newPassword == "" {
			l.Debug("Bad formdata", "Can't get old or new password")
			response = suckhttp.NewResponse(404, "Bad Request")
			return
		}

		if err := s.updatePassword(l, uid, oldPassword, newPassword); err != nil {
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
