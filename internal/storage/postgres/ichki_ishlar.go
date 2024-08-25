package postgres

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type IIRepo struct {
	db *sqlx.DB
}

func NewIIRepo(db *sqlx.DB) storage.IIStorage {
	return &IIRepo{
		db: db,
	}
}

// CreateFines creates a new fine record
func (i *IIRepo) CreateFines(fine *models.FineReq) error {
	query := `
		INSERT INTO fines (tech_passport_number, license_plate, officer_id, fine_owner, fine_date, price)
		VALUES ($1, $2, $3, $4, $5, $6)`
	err := i.db.QueryRowx(query, fine.TechPassportNumber, fine.LicensePlate, fine.OfficerID, fine.FineOwner, time.Now(), fine.Price)
	if err != nil {
		log.Println("Error creating fine:", err)
		return err.Err()
	}
	return nil
}

// AcceptFinesById updates the payment date of a fine to mark it as paid
func (i *IIRepo) AcceptFinesById(accept models.FineAccept) error {
	query := `UPDATE fines SET payment_date = $1 WHERE id = $2`
	_, err := i.db.Exec(query, accept.PaymentDate, accept.Id)
	if err != nil {
		log.Println("Error accepting fine:", err)
		return err
	}
	return nil
}

// GetPaidFines retrieves all paid fines (with payment date set)
func (i *IIRepo) GetPaidFines(pagination models.Pagination) (*models.Fines, error) {
	var fines models.Fines
	offset := (pagination.Page - 1) * pagination.Limit
	query := `SELECT * FROM fines WHERE payment_date IS NOT NULL LIMIT $1 OFFSET $2`
	err := i.db.Select(&fines, query, pagination.Limit, offset)
	if err != nil {
		log.Println("Error getting paid fines:", err)
		return nil, err
	}
	return &fines, nil
}

// GetUnpaidFines retrieves all unpaid fines (with no payment date set)
func (i *IIRepo) GetUnpaidFines(pagination models.Pagination) (*models.Fines, error) {
	var fines models.Fines
	offset := (pagination.Page - 1) * pagination.Limit
	query := `SELECT * FROM fines WHERE payment_date IS NULL LIMIT $1 OFFSET $2`
	err := i.db.Select(&fines, query, pagination.Limit, offset)
	if err != nil {
		log.Println("Error getting unpaid fines:", err)
		return nil, err
	}
	return &fines, nil
}

// GetAllFines retrieves all fines
func (i *IIRepo) GetAllFines(pagination models.Pagination) (*models.Fines, error) {
	var fines models.Fines
	offset := (pagination.Page - 1) * pagination.Limit
	query := `SELECT * FROM fines LIMIT $1 OFFSET $2`
	err := i.db.Select(&fines, query, pagination.Limit, offset)
	if err != nil {
		log.Println("Error getting all fines:", err)
		return nil, err
	}
	return &fines, nil
}
