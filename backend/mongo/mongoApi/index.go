package mongoApi

import (
	"github.com/big-larry/mgo"
)

func EnsurePathIndex(bucket *mgo.GridFS) error {
	return EnsureIndexKey(bucket.Files, "path")
}

func EnsureIndexKey(collection *mgo.Collection, indexKey string) error {
	return collection.EnsureIndexKey(indexKey)
}
