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
	Info     []Info        `bson:"info,omitempty"   json:"info,omitempty"`
	TestId   bson.ObjectId `bson:"testId,omitempty" json:"testId,omitempty"`
}

type Info struct {
	Type string `bson:"type" json:"type"`
	Data string `bson:"data" json:"data"`
}
