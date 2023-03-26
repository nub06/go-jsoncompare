package util

import (
	"encoding/json"
	"os"
)

func ReadFromFile(filePath string) []byte {
	res, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return res
}

func ParseJson(data []byte) interface{} {

	var obj interface{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}

	return obj
}
