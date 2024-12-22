package reader_templates

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

func ReadJson(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetStmt(handlerName string) (map[string]interface{}, error) {
	filePathMap := map[string]string{}
	filePath, exist := filePathMap[handlerName]
	if !exist {
		log.Println("there is no such operation:", handlerName)
		return nil, errors.New("there is no such operation")
	}
	data, err := GetJson(filePath)
	if err != nil {
		log.Println("error read json:", err)
		return nil, err
	}

	return data, nil

}

func GetJson(filePath string) (map[string]interface{}, error) {
	data, err := ReadJson(filePath)
	if err != nil {
		log.Println("Ошибка чтения JSON:", err)
		return nil, err
	}
	return data, nil
}
