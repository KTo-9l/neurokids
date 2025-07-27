package models

import (
	"github.com/big-larry/mgo/bson"
)

type Course struct {
	CourseId bson.ObjectId `bson:"_id"     json:"courseId,omitempty"`
	Lessons  []Lesson      `bson:"lessons" json:"lessons"`
}

type Lesson struct {
	CourseId bson.ObjectId `bson:"courseId"         json:"courseId"`
	Type     string        `bson:"type"             json:"type"`
	Info     interface{}   `bson:"info,omitempty"   json:"info,omitempty"`
	TestId   interface{}   `bson:"testId,omitempty" json:"testId,omitempty"`
}
