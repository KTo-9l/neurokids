package mongoApi

import (
	"fmt"

	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func ListFilesForPath(bucket *mgo.GridFS, path []string) (gfsFiles []GridFSFile, err error) {
	// query := bucket.Files.Find(bson.M{"path": bson.M{"$all": path}})
	query := bucket.Files.Find(bson.M{"path": path[0]})
	err = query.All(&gfsFiles)
	if err != nil {
		return nil, err
	}

	return gfsFiles, nil
}

func ListAllFiles(bucket *mgo.GridFS) (gfsFiles []GridFSFile, err error) {
	query := bucket.Files.Find(bson.M{})
	err = query.All(&gfsFiles)
	if err != nil {
		return nil, err
	}

	return gfsFiles, nil
}

func ListAllCollection(collection *mgo.Collection) (coll []interface{}, err error) {
	q := collection.Find(nil)
	err = q.All(&coll)
	return
}

func ListTypifiedCollection[T any](collection *mgo.Collection) (coll []T, err error) {
	q := collection.Find(nil)
	err = q.All(&coll)
	return
}

func ListTypifiedCollectionWithSelector[T any](collection *mgo.Collection, selector interface{}) (coll []T, err error) {
	q := collection.Find(selector)
	err = q.All(&coll)
	return
}

func ListSortedCollection(collection *mgo.Collection, sort []string) (coll []interface{}, err error) {
	q := collection.Find(nil).Sort(sort...)
	err = q.All(&coll)
	return
}

func ListTypifiedSortedCollectionWithSelector[T any](collection *mgo.Collection, selector interface{}, sort []string) (coll []T, err error) {
	q := collection.Find(selector).Sort(sort...)
	err = q.All(&coll)
	return
}

func ListTypifiedSortedCollectionByFields[T any](collection *mgo.Collection, selector map[string]interface{}, sort []string) (coll []T, err error) {
	q := collection.Find(selector).Sort(sort...)
	err = q.All(&coll)
	return
}

func ListLimitedCollectionByFields[T any](collection *mgo.Collection, selector map[string]interface{}, sort []string, amount int) (coll []T, err error) {
	q := collection.Find(selector).Sort(sort...).Limit(amount)
	err = q.All(&coll)
	return
}

func FindById[T any](collection *mgo.Collection, id string) (obj T, err error) {
	err = collection.FindId(bson.ObjectIdHex(id)).One(&obj)
	return
}

func FindOne[T any](collection *mgo.Collection, selector map[string]interface{}) (obj T, err error) {
	err = collection.Find(selector).One(&obj)
	return
}

func GetAllChunks(bucket *mgo.GridFS) {
	var fileMeta struct {
		Id       interface{} `bson:"_id"`
		Filename string      `bson:"filename"`
	}

	bucket.Files.Find(bson.M{"filename": "./testFile.txt"}).One(&fileMeta)
	// bucket.Files.Find(bson.M{"filename": "./test.msi"}).One(&fileMeta)
	fmt.Printf("Id for %v is: %v\n", fileMeta.Filename, fileMeta.Id)

	query := bucket.Chunks.Find(bson.M{"files_id": fileMeta.Id})

	var results any
	query.One(&results)

	fmt.Println(results)
}
