package main

import (
	"bytes"
	"mime"
	"mime/multipart"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateFilePathById(l logger.Logger, rbody *suckhttp.Request) (ok bool, err error) {
	_, params, _ := mime.ParseMediaType(rbody.GetHeader("content-type"))

	r := bytes.NewReader(rbody.Body)
	mr := multipart.NewReader(r, params["boundary"])
	form, err := mr.ReadForm(100 << 20)
	if err != nil {
		l.Error("Error reading form", err)
		return false, err
	}

	path, ok := form.Value["path"]
	if !ok {
		return false, nil
	}

	id, ok := form.Value["id"]
	if !ok {
		return false, nil
	}

	toUpdate := bson.M{
		"$set": bson.M{
			"path": path,
		},
	}

	err = mongoApi.UpdateById(s.bucket.Files, id[0], toUpdate)
	if err != nil {
		l.Error("UpdateFile", err)
		return
	}

	return true, nil
}
