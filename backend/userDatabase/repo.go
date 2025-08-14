package database

import (
	"context"
)

func GetUser(r *Repo, user_id int) (user User, err error) {
	query := "select u.id,u.name,u.email,u.unsubscribe,p.perm from users_perms p inner join users u on u.id=p.user_id where p.user_id=$1"
	var perm *int
	err = r.GetDB().QueryRow(context.Background(), query, user_id).Scan(&user.Id, &user.Name, &user.Email, &user.Unsubscribe, &perm)
	if perm != nil {
		user.Perm = *perm
	}
	return
}
