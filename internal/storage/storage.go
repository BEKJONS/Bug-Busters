package storage

import "bug_busters/pkg/models"

type AuthStorage interface {
	Register(in models.RegisterRequest) error
	Login(in *models.LoginEmailRequest) (*models.LoginResponse, error)
	AddLicence(in *models.LicenceNumber) error
}

type IIStorage interface {
	CreateFines(fine *models.FineReq) error
	AcceptFinesById(accept models.FineAccept) error
	GetPaidFines(pagination models.Pagination) (models.Fines, error)
	GetUnpaidFines(pagination models.Pagination) (models.Fines, error)
	GetAllFines(pagination models.Pagination) (*models.Fines, error)
}

type ServiceStorage interface {
	CreateService(service *models.Service) (*models.Service, error)
	GetService(id string) (*models.Service, error)
	UpdateService(service *models.Service) (*models.Service, error)
	DeleteService(id string) (string, error)
	GetServices() (*models.Services, error)
}
