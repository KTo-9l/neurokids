package main

import (
	"authorization/database"
	"context"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getAllUsers(l logger.Logger) (users []database.User, err error) {
	rows, err := s.rep.GetDB().Query(context.Background(), "SELECT id,name FROM users WHERE done = true")
	if err != nil {
		l.Error("Select query", err)
		return
	}

	for rows.Next() {
		var user database.User
		if err = rows.Scan(&user.Id, &user.Name); err != nil {
			l.Error("Scan user", err)
			return
		}

		users = append(users, user)
	}
	return
}
