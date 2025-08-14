package models

import "github.com/big-larry/mgo/bson"

type CourseUser struct {
	Id        bson.ObjectId `bson:"_id"               json:"id,omitempty"`
	UserId    int           `bson:"userId"            json:"userId"`
	CourseId  bson.ObjectId `bson:"courseId"          json:"courseId"`
	Purchased bool          `bson:"purchased"         json:"purchased"`
	Progress  Progress      `bson:"progress"          json:"progress"`
}

type Progress struct {
	Opened   bool `bson:"opened"   json:"opened"`
	Stage    int  `bson:"stage"    json:"stage"`
	Finished bool `bson:"finished" json:"finished"`
}
