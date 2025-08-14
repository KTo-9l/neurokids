package main

import (
	"authorization/database"
	"context"

	"github.com/google/uuid"
)

type UserPassword string
type UserLogin string

func AuthenticationUser(rep *database.Repo, login UserLogin, pass UserPassword) (string, int, string, error) {
	user_name := ""
	user_id := 0
	if err := rep.GetDB().QueryRow(context.Background(), "SELECT id,name FROM users where (login=$1 OR email=$1) AND password=$2 AND unsubscribe=false AND done=true", login, pass).Scan(&user_id, &user_name); err != nil {
		return "", 0, "", err
	}
	token := uuid.NewString()
	if _, err := rep.GetDB().Exec(context.Background(), "UPDATE users SET token=$3, last_login=NOW() WHERE (login=$1 OR email=$1) AND password=$2", login, pass, token); err != nil {
		return "", 0, "", err
	}
	return token, user_id, user_name, nil
}
