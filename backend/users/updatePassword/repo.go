package main

import (
	"context"
	"errors"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updatePassword(l logger.Logger, uid int, oldPassword, newPassword string) (err error) {
	var currentPassword string
	err = s.rep.GetDB().QueryRow(context.Background(), "SELECT password FROM users WHERE id=$1", uid).Scan(&currentPassword)
	if err != nil {
		l.Error("Select query", err)
		return
	}

	if currentPassword != oldPassword {
		err = errors.New("incorrent old password")
		return
	}

	_, err = s.rep.GetDB().Exec(context.Background(), "UPDATE users SET password = $1 WHERE id=$2", newPassword, uid)
	if err != nil {
		l.Error("Update query", err)
		return
	}

	return
}
