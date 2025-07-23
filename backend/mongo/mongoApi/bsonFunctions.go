package mongoApi

import (
	"github.com/big-larry/mgo/bson"
)

func ObjectIdToString(objectId bson.ObjectId) string {
	return objectId.Hex()
}

func ObjectIdInterfaceToString(inter interface{}) string {
	if objId, ok := inter.(bson.ObjectId); ok {
		return ObjectIdToString(objId)
	}
	return ""
}

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

func IsObjectId(id string) bool {
	return bson.IsObjectIdHex(id)
}
