package services

type GroupService struct {
	*GenericService
}

func NewHostService() *GroupService {
	return &GroupService{NewGenericService("host")}
}
