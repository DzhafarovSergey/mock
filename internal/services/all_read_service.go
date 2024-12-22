package services

type HostService struct {
	*GenericService
}

func NewHostService() *HostService {
	return &HostService{NewGenericService("host")}
}
