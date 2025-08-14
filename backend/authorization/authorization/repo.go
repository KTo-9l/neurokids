package main

import (
	"authorization/auth_helpers"
	"authorization/database"
	"context"
)

type UserId int
type UserToken string

// Тут было так, чторазные прова на разные данные, но я убрал - это лишнее пока... Поэтому ключ всегда один - 0
func AuthorizationUser(rep *database.Repo, token UserToken, user_id UserId) (map[int]auth_helpers.Perms, error) {
	row, err := rep.GetDB().Query(context.Background(), "SELECT perm,u.token FROM users_perms up INNER JOIN users u ON u.id=up.user_id WHERE u.id=$1 AND u.token=$2", user_id, token)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	result := make(map[int]auth_helpers.Perms)
	for row.Next() {
		var (
			perm auth_helpers.Perms
			t    string
		)
		if err = row.Scan(&perm, &t); err != nil {
			return nil, err
		}
		result[0] = perm
	}
	return result, nil
}
