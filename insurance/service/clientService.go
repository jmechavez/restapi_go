package service

import "github.com/jmechavez/restapi_go/insurance/domain"

type ClientService interface {
	GetAllClient() ([]domain.Client, error)
}

type DefaultClientService struct {
	repo domain.ClientRepository
}

func (r DefaultClientService) GetAllClient() ([]domain.Client, error) {
	return r.repo.FindAll()
}

func NewClientService(repository domain.ClientRepository) DefaultClientService {
	return DefaultClientService{repository}
}
