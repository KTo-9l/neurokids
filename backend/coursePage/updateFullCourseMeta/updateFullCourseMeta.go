package main

import (
	"encoding/json"

	"coursePage/models"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateFullCourseMeta(l logger.Logger, reqBytes []byte) (course models.CourseMeta, err error) {
	err = json.Unmarshal(reqBytes, &course)
	if err != nil {
		l.Error("Course Unmarshal", err)
		return models.CourseMeta{}, err
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"title":   course.Title,
			"author":  course.Author,
			"shown":   course.Shown,
			"deleted": course.Deleted,

			"shortCard": bson.M{
				"_id":         course.Id,
				"description": course.ShortCard.Description,
				"cover":       course.ShortCard.Cover,
				"city":        course.ShortCard.City,
				"time":        course.ShortCard.Time,
				"type":        course.ShortCard.Type,
			},

			"fullCard": bson.M{
				"_id":               course.Id,
				"headDescription":   course.FullCard.HeadDescription,
				"educationForm":     course.FullCard.EducationForm,
				"endDocument":       course.FullCard.EndDocument,
				"availableMaterial": course.FullCard.AvailableMaterial,
				"knowledgeBlock":    course.FullCard.KnowledgeBlock,
				"forComfortBlock":   course.FullCard.ForComfortBlock,
				"faq":               course.FullCard.Faq,
			},
		},
	}

	err = mongoApi.UpdateById(s.collection, course.Id, toUpdate)
	if err != nil {
		l.Error("coll.Update", err)
	}
	return
}
