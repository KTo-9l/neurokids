package database

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db  *pgxpool.Pool
	mux sync.Mutex
}

func OpenRepository(connectionString string) (*Repo, error) {
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return &Repo{db: db}, nil
}

func (rep *Repo) Close() {
	rep.db.Close()
}

func (rep *Repo) GetDB() *pgxpool.Pool {
	return rep.db
}

func (rep *Repo) Lock() {
	rep.mux.Lock()
}
func (rep *Repo) Unlock() {
	rep.mux.Unlock()
}
