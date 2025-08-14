package mongoApi

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func UpdateById(collection *mgo.Collection, id, newObject interface{}) (err error) {
	var objId bson.ObjectId
	if _, ok := id.(bson.ObjectId); ok {
		objId = id.(bson.ObjectId)
	} else if bson.IsObjectIdHex(id.(string)) {
		objId = bson.ObjectIdHex(id.(string))
	} else {
		return errors.New("not correct id")
	}

	return collection.UpdateId(objId, newObject)
}

func UpdateGridFSById(bucket *mgo.GridFS, id string, file *os.File, path []string) (ok bool, err error) {
	if !bson.IsObjectIdHex(id) {
		return false, nil
	}

	objId := bson.ObjectIdHex(id)

	err = bucket.RemoveId(objId)
	if err != nil {
		return false, err
	}

	_, err = InsertInGridFSWithId(bucket, file, path, objId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateGridFSByIdFromMultipart(bucket *mgo.GridFS, id string, fileHeader *multipart.FileHeader, path []string) (ok bool, err error) {
	if !bson.IsObjectIdHex(id) {
		return false, nil
	}

	objId := bson.ObjectIdHex(id)

	err = bucket.RemoveId(objId)
	if err != nil {
		return false, err
	}

	path = append(path, fileHeader.Filename)
	_, err = InsertInGridFSWithIdFromMultipart(bucket, fileHeader, path, objId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpsertGridFSByPathFromMultipart(bucket *mgo.GridFS, fileHeader *multipart.FileHeader, path []string) (ok bool, err error) {
	path = append(path, fileHeader.Filename)
	err = bucket.RemovePath(path)
	if err != nil {
		return false, err
	}

	_, err = InsertInGridFSFromMultipart(bucket, fileHeader, path)
	if err != nil {
		return false, err
	}

	return true, nil
}
