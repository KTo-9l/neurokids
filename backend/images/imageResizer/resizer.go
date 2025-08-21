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

func (s *service) getResizedImage(l logger.Logger, req *suckhttp.Request) (imageBytes []byte, err error) {
	pathFragments := strings.Split(req.Uri.Path[1:], "/")
	reqWidth, err := strconv.Atoi(pathFragments[1])
	if err != nil {
		l.Error("Parsing width (string) to (int)", err)
		return nil, err
	}
	if reqWidth > int(s.limitSize) {
		reqWidth = int(s.limitSize)
	}

	imagePath := strings.Join(pathFragments[2:len(pathFragments)-1], "/")

	imageName := pathFragments[len(pathFragments)-1]
	originalImagePath := fmt.Sprintf("/%s/%s", imagePath, imageName)

	originalImageBytes, err := s.getOriginalImage(l, originalImagePath)
	if err != nil {
		l.Error("GetOriginalImage", err)
		return nil, err
	}

	imageBytes, sourceWidth, err := s.resizeImage(l, originalImageBytes, reqWidth)
	err = s.saveImageInCache(l, imagePath, imageName, reqWidth, sourceWidth, imageBytes)

	return
}

func (s *service) updateImage(l logger.Logger, req *suckhttp.Request) (err error) {
	pathFragments := strings.Split(req.Uri.Path[1:], "/")
	imagePath := strings.Join(pathFragments[1:len(pathFragments)-1], "/")
	imageName := pathFragments[len(pathFragments)-1]

	originalImagePath := fmt.Sprintf("/%s/%s", imagePath, imageName)
	originalImageBytes, err := s.getOriginalImage(l, originalImagePath)
	if err != nil {
		l.Error("Error get original image", err)
		return err
	}

	dirEntries, err := os.ReadDir(s.cachedImagePath)
	for _, dir := range dirEntries {
		if !dir.IsDir() {
			continue
		}

		replacingImageDir := fmt.Sprintf("%s%s/%s", s.cachedImagePath, dir.Name(), strings.Join(pathFragments[1:], "/"))

		_, err = os.Open(replacingImageDir)
		if err != nil {
			l.Error("Opening file", err)
			continue
		}

		reqWidth, err := strconv.Atoi(dir.Name())
		if err != nil {
			continue
		}

		newImageBytes, sourceWidth, err := s.resizeImage(l, originalImageBytes, reqWidth)
		if err != nil {
			l.Error("Resize Image error", err)
			return err
		}
		err = s.saveImageInCache(l, imagePath, imageName, reqWidth, sourceWidth, newImageBytes)
		if err != nil {
			l.Error("Saving image in cache", err)
			return err
		}
	}
	return nil
}

func (s *service) resizeImage(l logger.Logger, originalImageBytes []byte, reqWidth int) (imgBytes []byte, sourceWidth int, err error) {
	originalImageReader := bytes.NewReader(originalImageBytes)

	img, format, err := image.Decode(originalImageReader)
	if err != nil {
		l.Debug("Format", format)
		l.Error("Decoding image", err)
		return nil, -1, err
	}

	sourceWidth = img.Bounds().Dx()
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
		return nil, -1, err
	}
	if err != nil {
		l.Error("Encoding image", err)
		return nil, -1, err
	}

	imgBytes = buf.Bytes()
	return imgBytes, sourceWidth, nil
}

func (s *service) saveImageInCache(l logger.Logger, imagePath, imageName string, reqWidth, sourceWidth int, imageBytes []byte) error {
	if reqWidth < sourceWidth {
		l.Info("Create Image", imageName)
		newImgDir := fmt.Sprintf("%s%d/%s", s.cachedImagePath, reqWidth, imagePath)
		newImgName := fmt.Sprintf("%s/%s", newImgDir, imageName)

		err := os.MkdirAll(newImgDir, 0755)
		if err != nil {
			l.Error("Creating image directory", err)
			return err
		}

		err = os.WriteFile(newImgName, imageBytes, 0665)
		if err != nil {
			l.Error("Creating image", err)
			return err
		}
	}
	return nil
}
