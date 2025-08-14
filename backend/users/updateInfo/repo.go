package main

import (
	"context"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateInfo(l logger.Logger, uid int, name, email string) (err error) {
	_, err = s.rep.GetDB().Exec(context.Background(), "UPDATE users SET name = $1, email = $2 WHERE id=$3", name, email, uid)
	if err != nil {
		l.Error("Update query", err)
		return
	}

	return
}
