package models

import "github.com/big-larry/mgo/bson"

type Test struct {
	Id        bson.ObjectId `bson:"_id"                   json:"id,omitempty"`
	Title     string        `bson:"title"                 json:"title"`
	Questions []Question    `bson:"questions"             json:"questions"`
	Deleted   bool          `bson:"deleted,omitempty"     json:"deleted,omitempty"`
}

type Question struct {
	Type     string        `bson:"type"               json:"type"`
	Info     interface{}   `bson:"info"               json:"info"`
	Variants []interface{} `bson:"variants,omitempty" json:"variants,omitempty"`
	Answer   interface{}   `bson:"answer"             json:"answer"`
}
