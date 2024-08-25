package storage

import "bug_busters/pkg/models"

type AuthStorage interface {
	Register(in models.RegisterRequest) error
	Login(in *models.LoginEmailRequest) (*models.LoginResponse, error)
	AddLicence(in *models.LicenceNumber) error
}
