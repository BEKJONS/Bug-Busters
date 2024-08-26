package service

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"log/slog"
)

type SWService interface {
	CreateLicense(in models.DriverLicense) (models.Message, error)
	CreatePassport(in models.CardId) (models.Message, error)
	GetLicenseAll() ([]models.DriverLicense, error)
	GetPassportAll() (models.PassportsId, error)
	DeleteLicense(in models.LicenceNumbers) (models.Message, error)
	DeletePassport(in models.PassportId) (models.Message, error)
}

func NewSWStorage(st storage.SWStorage, log *slog.Logger) SWService {
	return &singleWindowService{log, st}
}

type singleWindowService struct {
	log *slog.Logger
	st  storage.SWStorage
}

func (s *singleWindowService) CreateLicense(in models.DriverLicense) (models.Message, error) {
	err := s.st.CreateLicense(in)
	if err != nil {
		s.log.Error("Failed to create license", "error", err)
		return models.Message{}, err
	}

	return models.Message{"License created"}, nil
}

func (s *singleWindowService) CreatePassport(in models.CardId) (models.Message, error) {
	err := s.st.CreatePassport(in.ID)
	if err != nil {
		s.log.Error("Failed to create passport", "error", err)
		return models.Message{}, err
	}

	return models.Message{"Passport created"}, nil
}

func (s *singleWindowService) GetLicenseAll() ([]models.DriverLicense, error) {
	res, err := s.st.GetLicenseAll()
	if err != nil {
		s.log.Error("Failed to get license", "error", err)
		return []models.DriverLicense{}, err
	}

	return res, nil
}

func (s *singleWindowService) GetPassportAll() (models.PassportsId, error) {
	res, err := s.st.GetPassportAll()
	if err != nil {
		s.log.Error("Failed to get passports", "error", err)
		return models.PassportsId{}, err
	}

	return models.PassportsId{res}, nil
}

func (s *singleWindowService) DeleteLicense(in models.LicenceNumbers) (models.Message, error) {
	err := s.st.DeleteLicense(in.LicenceNumber)
	if err != nil {
		s.log.Error("Failed to delete license", "error", err)
		return models.Message{}, err
	}

	return models.Message{"License deleted"}, nil
}

func (s *singleWindowService) DeletePassport(in models.PassportId) (models.Message, error) {
	err := s.st.DeletePassport(in.PassportId)
	if err != nil {
		s.log.Error("Failed to delete passport", "error", err)
		return models.Message{}, err
	}

	return models.Message{"Passport deleted"}, nil
}
