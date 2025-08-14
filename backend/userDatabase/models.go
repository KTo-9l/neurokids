package database

type User struct {
	Id          int
	Name        string
	Email       string
	Unsubscribe bool
	Perm        int
}
