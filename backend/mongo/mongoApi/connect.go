package mongoApi

import (
	"github.com/big-larry/mgo"
)

func Connect(connectionString string) (*mgo.Session, error) {
	session, err := mgo.Dial(connectionString)
	return session, err
}
