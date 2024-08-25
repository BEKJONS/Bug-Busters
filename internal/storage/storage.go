package storage

import "bug_busters/pkg/models"

type AuthStorage interface {
	Register(in models.RegisterRequest) (models.RegisterResponse, error)
	LoginEmail(in models.LoginEmailRequest) (models.LoginResponse, error)
}
