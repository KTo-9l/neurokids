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

	return gfsFile, err
}

func GetFileByPath(bucket *mgo.GridFS, path []string) (gfsFile *mgo.GridFile, err error) {
	gfsFile, err = bucket.OpenPath(path)

	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	return gfsFile, err
}

func ListFfilesForPath(bucket *mgo.GridFS, path []string) (gfsFiles []GridFSFile, err error) {
	// query := bucket.Files.Find(bson.M{"path": bson.M{"$all": path}})
	query := bucket.Files.Find(bson.M{"path": path[0]})
	err = query.All(&gfsFiles)
	if err != nil {
		return nil, err
	}

	return gfsFiles, nil
}
