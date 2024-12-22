package services

import (
	"log"
	"mock/internal/reader_templates"
)

type GenericService struct {
	handlerName string
}

func NewGenericService(handlerName string) *GenericService {
	return &GenericService{handlerName: handlerName}
}

func (s *GenericService) List() map[string]interface{} {
	data, err := reader_templates.GetStmt(s.handlerName)
	if err != nil {
		log.Println("Ошибка получения данных для", s.handlerName, ":", err)
		return nil
	}
	return data
}
