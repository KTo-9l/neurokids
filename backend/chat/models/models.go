package models

import (
	"time"

	"github.com/big-larry/mgo/bson"
)

type Chat struct {
	Id      bson.ObjectId `bson:"_id"     json:"id,omitempty"`
	Title   string        `bson:"title"   json:"title"`
	Members []int         `bson:"members" json:"members"`
	IsGroup bool          `bson:"isGroup" json:"isGroup"`
	Deleted bool          `bson:"deleted" json:"deleted,omitempty"`
}

type Message struct {
	Id          bson.ObjectId `bson:"_id"                   json:"id,omitempty"`
	Chat        bson.ObjectId `bson:"chatId"                json:"chatId"`
	From        int           `bson:"from"                  json:"from"`
	Time        time.Time     `bson:"time"                  json:"time"`
	Text        string        `bson:"text,omitempty"        json:"text,omitempty"`
	Attachments []interface{} `bson:"attachments,omitempty" json:"attachments,omitmepty"`
	Deleted     bool          `bson:"deleted"               json:"deleted,omitempty"`
}
