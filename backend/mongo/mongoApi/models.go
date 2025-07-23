package mongoApi

import (
	"time"

	"github.com/big-larry/mgo/bson"
)

type GridFSFile struct {
	Id          interface{} `bson:"_id"`
	Filename    string      `bson:"filename"`
	Path        []string    `bson:"path"`
	UploadDate  time.Time   `bson:"uploadDate"`
	MD5         string      `bson:"md5"`
	ChunkSize   int         `bson:"chunkSize"`
	Length      int64       `bson:",minsize"`
	ContentType *string     `bson:"contentType,omitempty"`
	Metadata    *bson.Raw   `bson:",omitempty"`
}

// type gfsFile struct {
// 	Id          interface{} `bson:"_id"`
// 	ChunkSize   int         `bson:"chunkSize"`
// 	UploadDate  time.Time   `bson:"uploadDate"`
// 	Length      int64       `bson:",minsize"`
// 	MD5         string
// 	Filename    string    `bson:",omitempty"`
// 	ContentType string    `bson:"contentType,omitempty"`
// 	Metadata    *bson.Raw `bson:",omitempty"`
// }
