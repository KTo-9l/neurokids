package main

import (
	"bytes"
	"mime"
	"mime/multipart"

	"dikobra3/mongoApi"

	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) upsertFileByPath(l logger.Logger, rbody *suckhttp.Request) (ok bool, err error) {
	_, params, _ := mime.ParseMediaType(rbody.GetHeader("content-type"))

	r := bytes.NewReader(rbody.Body)
	mr := multipart.NewReader(r, params["boundary"])
	form, err := mr.ReadForm(100 << 20)
	if err != nil {
		l.Error("Error reading form", err)
		return false, err
	}

	files, ok := form.File["file"]
	if !ok {
		l.Debug("Files", "Can't get from form")
		return false, nil
	}

	path, ok := form.Value["path"]
	if !ok {
		l.Debug("Path", "Can't get from form")
		return false, nil
	}

	ok, err = mongoApi.UpsertGridFSByPathFromMultipart(s.bucket, files[0], path)
	if err != nil {
		l.Error("UpsertFile", err)
		return
	}

	return true, nil
}
