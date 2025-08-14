package main

import (
	"authorization/database"
	"context"
)

func LogoutUser(rep *database.Repo, uid int) error {
	if _, err := rep.GetDB().Exec(context.Background(), "UPDATE users SET token=NULL, last_login=NOW() WHERE id=$1", uid); err != nil {
		return err
	}
	return nil
}
