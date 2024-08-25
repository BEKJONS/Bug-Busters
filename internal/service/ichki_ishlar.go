package service

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"log/slog"
)

type IIService interface {
	CreateFines(fine *models.FineReq) error
	AcceptFinesById(accept models.FineAccept) error
	GetPaidFines(pagination models.Pagination) (*models.Fines, error)
	GetUnpaidFines(pagination models.Pagination) (*models.Fines, error)
	GetAllFines(pagination models.Pagination) (*models.Fines, error)
}

func NewIIService(st storage.IIStorage, logger *slog.Logger) IIService {
	return &iiService{st, logger}
}

type iiService struct {
	st  storage.IIStorage
	log *slog.Logger
}

// CreateFines creates a new fine
func (s *iiService) CreateFines(fine *models.FineReq) error {
	s.log.Info("CreateFines started")
	err := s.st.CreateFines(fine)
	if err != nil {
		s.log.Error("failed to create fine", "error", err)
		return err
	}
	s.log.Info("fine created successfully")
	return nil
}

// AcceptFinesById accepts a fine by updating its payment date
func (s *iiService) AcceptFinesById(accept models.FineAccept) error {
	s.log.Info("AcceptFinesById started")
	err := s.st.AcceptFinesById(accept)
	if err != nil {
		s.log.Error("failed to accept fine", "error", err)
		return err
	}
	s.log.Info("fine accepted successfully", "fine_id", accept.Id)
	return nil
}

// GetPaidFines retrieves all paid fines
func (s *iiService) GetPaidFines(pagination models.Pagination) (*models.Fines, error) {
	s.log.Info("GetPaidFines started")
	fines, err := s.st.GetPaidFines(pagination)
	if err != nil {
		s.log.Error("failed to get paid fines", "error", err)
		return nil, err
	}
	s.log.Info("GetPaidFines completed successfully")
	return fines, nil
}

// GetUnpaidFines retrieves all unpaid fines
func (s *iiService) GetUnpaidFines(pagination models.Pagination) (*models.Fines, error) {
	s.log.Info("GetUnpaidFines started")
	fines, err := s.st.GetUnpaidFines(pagination)
	if err != nil {
		s.log.Error("failed to get unpaid fines", "error", err)
		return nil, err
	}
	s.log.Info("GetUnpaidFines completed successfully")
	return fines, nil
}

// GetAllFines retrieves all fines
func (s *iiService) GetAllFines(pagination models.Pagination) (*models.Fines, error) {
	s.log.Info("GetAllFines started")
	fines, err := s.st.GetAllFines(pagination)
	if err != nil {
		s.log.Error("failed to get all fines", "error", err)
		return nil, err
	}
	s.log.Info("GetAllFines completed successfully")
	return fines, nil
}
