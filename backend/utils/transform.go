package utils

import "encoding/json"

func ObjectToBytes(obj interface{}) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	return bytes, err
}
