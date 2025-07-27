package models

import (
	"time"

	"github.com/big-larry/mgo/bson"
)

type CourseMeta struct {
	Id        bson.ObjectId `bson:"_id"       json:"id,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	Title     string        `bson:"title"     json:"title"`
	Author    string        `bson:"author"    json:"author,omitempty"`
	ShortCard ShortCard     `bson:"shortCard" json:"shortCard"`
	FullCard  FullCard      `bson:"fullCard"  json:"fullCard"`
	Shown     bool          `bson:"shown"     json:"shown"`
	Deleted   bool          `bson:"deleted"   json:"deleted,omitempty"`
}

type ShortCard struct {
	CourseId    bson.ObjectId `bson:"_id"            json:"courseId"`
	Description string        `bson:"description"    json:"description"`
	Cover       interface{}   `bson:"cover"          json:"cover"`
	City        string        `bson:"city,omitempty" json:"city,omitempty"`
	Time        string        `bson:"time,omitempty" json:"time,omitempty"`
	Type        string        `bson:"type,omitempty" json:"type,omitempty"`
}

type FullCard struct {
	CourseId          bson.ObjectId `bson:"_id" json:"courseId"`
	HeadDescription   string        `bson:"headDescription" json:"headDescription"`
	EducationForm     string        `bson:"educationForm"   json:"educationForm"`
	EndDocument       string        `bson:"endDocument"     json:"endDocument"`
	AvailableMaterial struct {
		Modules       uint `bson:"modules"       json:"modules"`
		TheoryHours   uint `bson:"theoryHours"   json:"theoryHours"`
		PracticeHours uint `bson:"practiceHours" json:"practiceHours"`
		Weeks         uint `bson:"weeks"         json:"weeks"`
	} `bson:"availableMaterial" json:"availableMaterial"`
	KnowledgeBlock struct {
		Header    string `bson:"header" json:"header"`
		Knowledge []struct {
			Head string `bson:"head" json:"head"`
			Body string `bson:"body" json:"body"`
		} `bson:"knowledge" json:"knowledge"`
	} `bson:"knowledgeBlock" json:"knowledgeBlock"`
	ForComfortBlock struct {
		Head       string `bson:"head" json:"head"`
		Instrument []struct {
			Icon interface{} `bson:"icon" json:"icon"`
			Text string      `bson:"text" json:"text"`
		} `bson:"instruments" json:"instruments"`
	} `bson:"forComfortBlock" json:"forComfortBlock"`
	Faq []struct {
		Question string `bson:"question" json:"question"`
		Answer   string `bson:"answer"   json:"answer"`
	} `bson:"faq" json:"faq"`
}
