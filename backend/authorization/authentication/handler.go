package main

import (
	"authorization/auth_helpers"
	"bytes"
	"fmt"
	"net/url"

	"github.com/big-larry/suckhttp"
	"github.com/jackc/pgx"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	response = suckhttp.NewResponse(500, "Server error")
	redirect := req.Uri.Query().Get("redirect")
	if redirect == "" {
		redirect = "/"
	}
	var login, pass = "", ""
	if req.GetMethod() == suckhttp.GET {
	} else if req.GetMethod() == suckhttp.POST {
		var form url.Values
		form, err = auth_helpers.ParseForm(req)
		if err != nil {
			response = suckhttp.NewResponse(500, "Server error")
			return
		}
		login, _ = auth_helpers.TryGetFormRawValue(form, "login")
		pass, _ = auth_helpers.TryGetFormRawValue(form, "pass")
		if login != "" || pass != "" {
			user_token, user_id, username, err1 := AuthenticationUser(s.rep, UserLogin(login), UserPassword(pass))
			if err1 != nil {
				if err1 == pgx.ErrNoRows {
					goto start
				}
				err = err1
				return
			}
			l.Info("login", fmt.Sprint(user_token, user_id, username))
			token, err1 := auth_helpers.CreateToken(user_id, user_token, username)
			if err1 != nil {
				err = err1
				return
			}

			response = suckhttp.NewResponse(302, "Redirect")
			// response = suckhttp.NewResponse(200, "Ok")
			response.AddHeader(suckhttp.Set_Cookie, "token="+token+";path=/;expires=-1")
			// response.AddHeader(suckhttp.Location, redirect)
			response.AddHeader(suckhttp.Location, "/setting")

			return
		}
	}
start:
	var w bytes.Buffer
	if err = s.index_template.Execute(&w, struct {
		Login string
		Pass  string
	}{login, pass}); err != nil {
		response = suckhttp.NewResponse(500, "Server error")
	} else {
		response = suckhttp.NewResponse(200, "Ok").SetBody(w.Bytes())
	}
	return
}
