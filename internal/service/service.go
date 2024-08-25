package service

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
)

type IService interface {
	GetService(id string) (*models.Service, error)
	CreateService(service *models.Service) (*models.Service, error)
	UpdateService(service *models.Service) (*models.Service, error)
	DeleteService(id string) (string, error)
	GetServices() (*models.Services, error)
}

func NewService(st storage.ServiceStorage) IService {
	return &service{storage: st}
}

type service struct {
	storage storage.ServiceStorage
}

func (s *service) GetService(id string) (*models.Service, error) {
	return s.storage.GetService(id)
}

func (s *service) CreateService(service *models.Service) (*models.Service, error) {
	return s.storage.CreateService(service)
}

func (s *service) UpdateService(service *models.Service) (*models.Service, error) {
	return s.storage.UpdateService(service)
}

func (s *service) DeleteService(id string) (string, error) {
	return s.storage.DeleteService(id)
}

func (s *service) GetServices() (*models.Services, error) {
	return s.storage.GetServices()
}
