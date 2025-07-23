package models

import (
	"time"

	"github.com/big-larry/mgo/bson"
)

type Course struct {
	Id        bson.ObjectId `bson:"_id"               json:"id,omitempty"`
	Title     string        `bson:"title"             json:"title"`
	Lessons   []Lesson      `bson:"lessons"           json:"lessons"`
	CreatedAt time.Time     `bson:"createdAt"         json:"createdAt"`
	Author    string        `bson:"author,omitempty"  json:"author,omitempty"`
	Deleted   bool          `bson:"deleted,omitempty" json:"deleted,omitempty"`
}

type Lesson struct {
	Type   string      `bson:"type"             json:"type"`
	Info   interface{} `bson:"info,omitempty"   json:"info,omitempty"`
	TestId interface{} `bson:"testId,omitempty" json:"testId,omitempty"`
}

type CourseCard struct {
	CourseId   bson.ObjectId `bson:"courseId" json:"courseId"`
	InfoBlocks []interface{} `bson:"infoBlocks" json:"infoBlocks"`
}
