package main

import (
	"bytes"
	"mime"
	"mime/multipart"

	"dikobra3/mongoApi"

	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) updateFile(l logger.Logger, rbody *suckhttp.Request) (ok bool, err error) {
	_, params, _ := mime.ParseMediaType(rbody.GetHeader("content-type"))

	r := bytes.NewReader(rbody.Body)
	mr := multipart.NewReader(r, params["boundary"])
	form, err := mr.ReadForm(100 << 20)
	if err != nil {
		l.Error("Error reading form", err)
		return false, err
	}

	files, ok := form.File["files"]
	if !ok {
		return false, nil
	}

	path, ok := form.Value["path"]
	if !ok {
		return false, nil
	}

	id, ok := form.Value["id"]
	if !ok {
		return false, nil
	}

	ok, err = mongoApi.UpdateGridFSByIdFromMultipart(s.bucket, id[0], files[0], path)
	if err != nil {
		l.Error("UpdateFile", err)
		return
	}

	return true, nil
}
