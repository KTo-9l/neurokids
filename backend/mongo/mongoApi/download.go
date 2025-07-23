package mongoApi

import (
	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func GetFileById(bucket *mgo.GridFS, idString string) (gfsFile *mgo.GridFile, err error) {
	if bson.IsObjectIdHex(idString) {
		idHex := bson.ObjectIdHex(idString)
		gfsFile, err = bucket.OpenId(idHex)
	} else {
		gfsFile, err = bucket.OpenId(idString)
	}

	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	return gfsFile, nil
}
