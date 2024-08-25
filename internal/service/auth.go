package service

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/hashing"
	"bug_busters/pkg/models"
	"bug_busters/pkg/token"
	"errors"
	"log/slog"
)

type AuthService interface {
	Register(in models.RegisterRequest) error
	Login(in *models.LoginEmailRequest) (*models.Tokens, error)
}

func NewAuthService(st storage.AuthStorage, logger *slog.Logger) AuthService {
	return &authService{st, logger}
}

type authService struct {
	st  storage.AuthStorage
	log *slog.Logger
}

func (a *authService) Register(in models.RegisterRequest) error {
	hash, err := hashing.HashPassword(in.Password)
	if err != nil {
		a.log.Error("Failed to hash password", "error", err)
		return err
	}

	in.Password = hash

	err = a.st.Register(in)
	if err != nil {
		a.log.Error("Failed to register user", "error", err)
		return err
	}

	return nil
}

func (a *authService) Login(in *models.LoginEmailRequest) (*models.Tokens, error) {
	res, err := a.st.Login(in)
	if err != nil {
		a.log.Error("Failed to login", "error", err)
		return nil, err
	}

	check := hashing.CheckPasswordHash(res.Password, in.Password)
	if !check {
		a.log.Error("Invalid password")
		return nil, errors.New("invalid password")
	}

	refreshToken, err := token.GenerateRefreshToken(res.Id, res.Role, res.Password)
	if err != nil {
		a.log.Error("Failed to generate refresh token", "error", err)
		return nil, err
	}

	accessToken, err := token.GenerateAccessToken(res.Id, res.Role, res.Password)
	if err != nil {
		a.log.Error("Failed to generate access token", "error", err)
		return nil, err
	}

	response := &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Data:         *res,
	}

	return response, nil
}
