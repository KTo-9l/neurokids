package mongoApi

import (
	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func RemoveById(collection *mgo.Collection, id string) error {
	return collection.RemoveId(bson.ObjectIdHex(id))
}

func Remove(collection *mgo.Collection, selector map[string]interface{}) error {
	return collection.Remove(selector)
}

func RemoveAll(collection *mgo.Collection, selector map[string]interface{}) error {
	_, err := collection.RemoveAll(selector)
	return err
}
