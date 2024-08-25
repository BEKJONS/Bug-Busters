package postgres

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"database/sql"
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
	log.Println("Hello world", fine.TechPassportNumber)
	log.Println("Hello world", fine.TechPassportNumber)
	log.Println("Hello world", fine.TechPassportNumber)

	query := `
		INSERT INTO fines (tech_passport_number, license_plate, officer_id, fine_owner, fine_date, price)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := i.db.Exec(query, fine.TechPassportNumber, fine.LicensePlate, fine.OfficerID, fine.FineOwner, time.Now(), fine.Price)
	if err != nil {
		log.Println("Error creating fine:", err)
		return err
	}
	return nil
}

// AcceptFinesById updates the payment date of a fine to mark it as paid
func (i *IIRepo) AcceptFinesById(accept models.FineAccept) error {
	query := `UPDATE fines SET payment_date = $1 WHERE id = $2`
	_, err := i.db.Exec(query, time.Now(), accept.Id)
	if err != nil {
		log.Println("Error accepting fine:", err)
		return err
	}
	return nil
}

// GetPaidFines retrieves all paid fines (with payment date set)
func (i *IIRepo) GetPaidFines(pagination models.Pagination) (models.Fines, error) {
	var fines models.Fines
	offset := (pagination.Page - 1) * pagination.Limit
	query := `SELECT * FROM fines WHERE payment_date IS NOT NULL LIMIT $1 OFFSET $2`
	err := i.db.Select(&fines, query, pagination.Limit, offset)
	if err != nil {
		log.Println("Error getting paid fines:", err)
		return nil, err
	}
	return fines, nil
}

// GetUnpaidFines retrieves all unpaid fines (with no payment date set)
func (i *IIRepo) GetUnpaidFines(pagination models.Pagination) (models.Fines, error) {
	offset := (pagination.Page - 1) * pagination.Limit
	query := `SELECT id, tech_passport_number, license_plate, officer_id, fine_owner, fine_date, price FROM fines WHERE payment_date IS NULL LIMIT $1 OFFSET $2`
	row, err := i.db.Query(query, pagination.Limit, offset)
	if err != nil {
		log.Println("Error getting unpaid fines:", err)
		return nil, err
	}
	var fineses []models.Fine
	for row.Next() {
		fine := models.Fine{}
		err = row.Scan(&fine.ID, &fine.TechPassportNumber, &fine.LicensePlate, &fine.OfficerID, &fine.FineOwner, &fine.FineDate, &fine.Price)
		fine.PaymentDate = ""
		if err != nil {
			log.Println("Error getting unpaid fines:", err)
			return nil, err
		}
		fineses = append(fineses, fine)
	}
	return fineses, nil
}

func (i *IIRepo) GetAllFines(pagination models.Pagination) (*models.Fines, error) {
	var fines models.Fines
	offset := (pagination.Page - 1) * pagination.Limit
	query := `SELECT * FROM fines LIMIT $1 OFFSET $2`
	row, err := i.db.Query(query, pagination.Limit, offset)
	if err != nil {
		log.Println("Error getting all fines:", err)
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		fine := models.Fine{}
		var paymentDate sql.NullTime
		err = row.Scan(&fine.ID, &fine.TechPassportNumber, &fine.LicensePlate, &fine.OfficerID, &fine.FineOwner, &fine.FineDate, &fine.Price, &paymentDate)

		if err != nil {
			log.Println("Error scanning fines:", err)
			return nil, err
		}

		// Convert sql.NullTime to string
		if paymentDate.Valid {
			fine.PaymentDate = paymentDate.Time.Format("2006-01-02 15:04:05")
		} else {
			fine.PaymentDate = ""
		}

		fines = append(fines, fine)
	}

	return &fines, nil
}

// ConvertNullTimeToString converts sql.NullTime to a string
func ConvertNullTimeToString(nt sql.NullTime) string {
	if nt.Valid {
		// Convert time.Time to string with a format of your choice
		return nt.Time.Format(time.RFC3339) // Example format
	}
	return "" // or some default value
}
