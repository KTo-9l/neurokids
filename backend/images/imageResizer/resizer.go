package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/big-larry/suckhttp"
	"github.com/nfnt/resize"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) getOriginalImage(l logger.Logger, path string) (imageBytes []byte, err error) {
	pub := s.pubs_getter.Get(filesharing_get_file_by_path)
	if pub == nil {
		return nil, errors.New("file server is shutdown")
	}

	uri := fmt.Sprintf("/?path=%s", path[1:])

	req, err := suckhttp.NewRequest("GET", uri)
	if err != nil {
		l.Error("Create Request For Original Image", err)
		return
	}

	resp, err := pub.SendHTTP(req)
	if err != nil {
		l.Error("Getting Response From Original Image", err)
		return
	}

	if statusCode, _ := resp.GetStatus(); statusCode != 200 {
		err = errors.New("not found image")
		l.Error("Not found image", err)
		return
	}

	return resp.GetBody(), nil
}

func (s *service) getResizedImage(l logger.Logger, req *suckhttp.Request) (imgBytes []byte, err error) {
	pathFragments := strings.Split(req.Uri.Path[1:], "/")
	reqWidth, err := strconv.Atoi(pathFragments[1])
	if err != nil {
		l.Error("Parsing width (string) to (int)", err)
		return nil, err
	}

	imagePath := strings.Join(pathFragments[2:len(pathFragments)-1], "/")
	l.Debug("ImagePath ==============", fmt.Sprint(imagePath))

	imageName := pathFragments[len(pathFragments)-1]
	originalImagePath := fmt.Sprintf("/%s/%s", imagePath, imageName)
	l.Debug("OriginalImagePath", originalImagePath)

	originalImageBytes, err := s.getOriginalImage(l, originalImagePath)
	if err != nil {
		l.Error("GetOriginalImage", err)
		return nil, err
	}

	originalImageReader := bytes.NewReader(originalImageBytes)

	img, format, err := image.Decode(originalImageReader)
	if err != nil {
		l.Debug("Format", format)
		l.Error("Decoding image", err)
		return nil, err
	}

	// sourceWidth := img.Bounds().Dx()
	img = resize.Resize(uint(reqWidth), 0, img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	switch format {
	case "png":
		err = png.Encode(buf, img)
	case "jpeg":
		err = jpeg.Encode(buf, img, nil)
	default:
		err = errors.New("invalid format requested")
		l.Error("Image format", err)
		return nil, err
	}
	if err != nil {
		l.Error("Encoding image", err)
		return nil, err
	}

	imgBytes = buf.Bytes()
	// if reqWidth < sourceWidth {
	// 	l.Info("Create Image", imageName)
	// 	newImgDir := fmt.Sprintf("./../../../static/images/%d/%s/%s", reqWidth, pathFragments[2], pathFragments[3])
	// 	newImgName := fmt.Sprintf("%s/%s", newImgDir, imageName)

	// 	err := os.MkdirAll(newImgDir, 0755) // 0755 — стандартные права доступа
	// 	if err != nil {
	// 		fmt.Printf("Failed to create directory: %v\n", err)
	// 		return nil, err
	// 	}

	// 	err = os.WriteFile(newImgName, imgBytes, 0665)
	// 	if err != nil {
	// 		fmt.Printf("Creating image: %v\n", err)
	// 		return nil, err
	// 	}
	// }

	return
}

func getNewImage(req *suckhttp.Request, lg logger.Logger) *suckhttp.Response {
	pathFragments := strings.Split(req.Uri.Path[8:], "/")
	reqWidth, err := strconv.Atoi(pathFragments[0])
	if err != nil {
		println("Error parse width (string) to (int)", err)
		return suckhttp.NewResponse(500, "Internal Server Error")
	}

	imageName := pathFragments[1]
	pathToSourceImage := fmt.Sprintf("./../../static/images/source/%s", imageName)
	defaultImageFile, err := os.Open(pathToSourceImage)
	if err != nil {
		println("Error opening file")
		return suckhttp.NewResponse(500, "Internal Server Error")
	}
	defer defaultImageFile.Close()

	format := getFileFormat(imageName)

	var img image.Image
	img, _, err = image.Decode(defaultImageFile)
	if err != nil {
		println("Error decoding image")
		return suckhttp.NewResponse(500, "Internal Server Error")
	}

	sourceWidth := img.Bounds().Dx()
	img = resize.Resize(uint(reqWidth), 0, img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	switch format {
	case "png":
		err = png.Encode(buf, img)
	case "jpeg":
		err = jpeg.Encode(buf, img, nil)
	default:
		println("Some error")
		return suckhttp.NewResponse(400, "Bad Request")
	}
	if err != nil {
		println("Encoding error", err)
		return suckhttp.NewResponse(500, "Internal Server Error")
	}

	imgBytes := buf.Bytes()
	if reqWidth < sourceWidth {
		lg.Info("Create Image", imageName)
		newImgName := fmt.Sprintf("./../../static/images/%d/%s", reqWidth, imageName)
		err = os.WriteFile(newImgName, imgBytes, 0662)
		if err != nil {
			println("Error creating image", err)
			return suckhttp.NewResponse(500, "Internal Server Error")
		}
	}
	return suckhttp.NewResponse(200, "Ok").SetBody(imgBytes)
}

func getFileFormat(path string) string {
	lower := strings.ToLower(path)
	if strings.Contains(lower, ".jpg") || strings.Contains(lower, ".jpeg") {
		return "jpeg"
	} else if strings.Contains(lower, ".png") {
		return "png"
	}
	return "unknown"
}
