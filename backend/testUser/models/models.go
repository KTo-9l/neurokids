package models

import "github.com/big-larry/mgo/bson"

type TestUser struct {
	Id       bson.ObjectId `bson:"_id"      json:"id,omitempty"`
	UserId   bson.ObjectId `bson:"userId"   json:"userId"`
	TestId   bson.ObjectId `bson:"testId"   json:"testId"`
	Progress Progress      `bson:"progress" json:"progress"`
}

type Progress struct {
	Opened   bool `bson:"opened"   json:"opened"`
	Stage    int  `bson:"stage"    json:"stage"`
	Finished bool `bson:"finished" json:"finished"`
	Correct  int  `bson:"correct"  json:"correct"`
}
