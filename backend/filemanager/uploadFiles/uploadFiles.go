package main

import (
	"bytes"
	"dikobra3/mongoApi"
	"mime"
	"mime/multipart"

	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) uploadFiles(l logger.Logger, rbody *suckhttp.Request) (filesIds interface{}, err error) {
	_, params, _ := mime.ParseMediaType(rbody.GetHeader("content-type"))

	r := bytes.NewReader(rbody.Body)
	mr := multipart.NewReader(r, params["boundary"])
	form, err := mr.ReadForm(100 << 20)
	if err != nil {
		l.Error("Error reading form", err)
		return nil, err
	}

	files, ok := form.File["files"]
	if !ok {
		return nil, nil
	}

	path, ok := form.Value["path"]
	if !ok {
		return nil, nil
	}

	var resp []struct {
		Filename string      `json:"filename"`
		Path     []string    `json:"path"`
		Id       interface{} `json:"id"`
	}
	for _, fileHeader := range files {
		resultPath := append(path, fileHeader.Filename)
		id, err := mongoApi.InsertInGridFSFromMultipart(s.bucket, fileHeader, resultPath)
		if err != nil {
			l.Error("UploadFile", err)
			return nil, err
		}
		resp = append(resp, struct {
			Filename string      `json:"filename"`
			Path     []string    `json:"path"`
			Id       interface{} `json:"id"`
		}{Filename: fileHeader.Filename, Path: resultPath, Id: id})
	}
	return resp, nil
}
