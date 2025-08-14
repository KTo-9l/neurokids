package mongoApi

import (
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/big-larry/mgo"
)

func InsertInGridFSWithId(bucket *mgo.GridFS, file *os.File, path []string, id interface{}) (interface{}, error) {
	filename := file.Name()
	path = append(path, filename)

	gridFSFile, err := bucket.Create(filename, path)
	if err != nil {
		log.Println("Error creating gfs file:", err)
		return nil, err
	}

	gridFSFile.SetId(id)

	_, err = io.Copy(gridFSFile, file)
	if err != nil {
		log.Println("Error copying file to gridFS:", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
		return nil, err
	}

	err = gridFSFile.Close()
	if err != nil {
		log.Println("Error closing gridFS File:", err)
		return nil, err
	}

	return gridFSFile.Id(), nil
}

func InsertInGridFS(bucket *mgo.GridFS, file *os.File, path []string) (interface{}, error) {
	filename := file.Name()
	path = append(path, filename)

	gridFSFile, err := bucket.Create(filename, path)
	if err != nil {
		log.Println("Error creating gfs file:", err)
		return nil, err
	}

	_, err = io.Copy(gridFSFile, file)
	if err != nil {
		log.Println("Error copying file to gridFS:", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
		return nil, err
	}

	err = gridFSFile.Close()
	if err != nil {
		log.Println("Error closing gridFS File:", err)
		return nil, err
	}

	return gridFSFile.Id(), nil
}

func InsertInGridFSWithIdFromMultipart(bucket *mgo.GridFS, fileHeader *multipart.FileHeader, path []string, id interface{}) (interface{}, error) {
	gridFSFile, err := bucket.Create(fileHeader.Filename, path)
	if err != nil {
		log.Println("Error creating gfs file:", err)
		return nil, err
	}

	gridFSFile.SetId(id)

	file, err := fileHeader.Open()
	_, err = io.Copy(gridFSFile, file)
	if err != nil {
		log.Println("Error copying file to gridFS:", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
		return nil, err
	}

	err = gridFSFile.Close()
	if err != nil {
		log.Println("Error closing gridFS File:", err)
		return nil, err
	}

	return gridFSFile.Id(), nil
}

func InsertInGridFSFromMultipart(bucket *mgo.GridFS, fileHeader *multipart.FileHeader, path []string) (interface{}, error) {
	gridFSFile, err := bucket.Create(fileHeader.Filename, path)
	if err != nil {
		log.Println("Error creating gfs file:", err)
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Error open file:", err)
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(gridFSFile, file)
	if err != nil {
		log.Println("Error copying file to gridFS:", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
		return nil, err
	}

	err = gridFSFile.Close()
	if err != nil {
		log.Println("Error closing gridFS File:", err)
		return nil, err
	}

	return gridFSFile.Id(), nil
}
