package mongoApi

import "github.com/big-larry/mgo"

func Insert(collection *mgo.Collection, docs interface{}) error {
	return collection.Insert(docs)
}

func InsertMany(collection *mgo.Collection, docs []interface{}) error {
	return collection.Insert(docs...)
}
