package service

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"log/slog"
)

type UserService interface {
	GetProfile(in models.UserId) (models.UserProfile, error)
	AddImage(in *models.UpdateCarImage) (models.Message, error)
	GetImage(in models.UserId) (models.Url, error)
	GetPaidFines(in models.UserId) (*[]*models.UserFines, error)
	GetUnpaid(in models.UserId) (*[]*models.UserFines, error)
	DeleteUser(in models.UserId) (models.Message, error)
}

func NewUserService(log *slog.Logger, st storage.UserStorage) UserService {
	return &userService{log, st}
}

type userService struct {
	log *slog.Logger
	st  storage.UserStorage
}

func (u *userService) GetProfile(in models.UserId) (models.UserProfile, error) {
	res, err := u.st.GetProfile(in)
	if err != nil {
		u.log.Error("Failed to get user profile", "error", err)
		return models.UserProfile{}, err
	}

	return res, nil
}

func (u *userService) AddImage(in *models.UpdateCarImage) (models.Message, error) {
	err := u.st.AddImage(in)
	if err != nil {
		u.log.Error("Failed to add image", "error", err)
		return models.Message{}, err
	}

	return models.Message{"Image added"}, nil
}

func (u *userService) GetImage(in models.UserId) (models.Url, error) {
	res, err := u.st.GetImage(in.Id)
	if err != nil {
		u.log.Error("Failed to get image", "error", err)
		return models.Url{}, err
	}

	return models.Url{Url: res}, nil
}

func (u *userService) GetPaidFines(in models.UserId) (*[]*models.UserFines, error) {
	res, err := u.st.GetPaidFines(in.Id)
	if err != nil {
		u.log.Error("Failed to get paid fines", "error", err)
		return nil, err
	}

	return res, nil
}

func (u *userService) GetUnpaid(in models.UserId) (*[]*models.UserFines, error) {
	res, err := u.st.GetUnpaid(in.Id)
	if err != nil {
		u.log.Error("Failed to get unpaid fines", "error", err)
		return nil, err
	}

	return res, nil
}

func (u *userService) DeleteUser(in models.UserId) (models.Message, error) {
	err := u.st.DeleteUser(in.Id)
	if err != nil {
		u.log.Error("Failed to delete user", "error", err)
		return models.Message{}, err
	}

	return models.Message{"User deleted"}, nil
}
