package auth_helpers

import "time"

type AuthResult struct {
	Token   string
	Uid     int
	Name    string
	Created time.Time
}

type AuthResultWithPerms[T ~int] struct {
	AuthResult
	Perms map[int]T
}
