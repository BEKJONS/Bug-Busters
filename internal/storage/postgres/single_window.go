package postgres

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"github.com/jmoiron/sqlx"
)

type SWRepo struct {
	db *sqlx.DB
}

// NewSWRepo creates a new SWRepo with a given database connection.
func NewSWRepo(db *sqlx.DB) storage.SWStorage {
	return &SWRepo{db: db}
}

// CreateLicense inserts a new driver license into the database.
func (repo *SWRepo) CreateLicense(license models.DriverLicense) error {
	query := `INSERT INTO driver_licenses (id, first_name, last_name, father_name, birth_date, address, issue_date, expiration_date, category, issued_by, license_number) 
              VALUES (gen_random_uuid(), :first_name, :last_name, :father_name, :birth_date, :address, :issue_date, :expiration_date, :category, :issued_by, :license_number)`

	_, err := repo.db.NamedExec(query, license)
	return err
}

// CreatePassport generates a tech passport UUID for a car in the database.
func (repo *SWRepo) CreatePassport(carID string) error {
	query := `UPDATE cars SET tech_passport_number = gen_random_uuid() WHERE id = $1`
	_, err := repo.db.Exec(query, carID)
	return err
}

// GetLicenseAll retrieves all driver licenses from the database.
func (repo *SWRepo) GetLicenseAll() ([]models.DriverLicense, error) {
	var licenses []models.DriverLicense
	query := `SELECT * FROM driver_licenses`

	err := repo.db.Select(&licenses, query)
	return licenses, err
}

// GetPassportAll retrieves all tech passport UUIDs from the database.
func (repo *SWRepo) GetPassportAll() ([]string, error) {
	var passports []string
	query := `SELECT tech_passport_number FROM cars`

	err := repo.db.Select(&passports, query)
	return passports, err
}

// DeleteLicense deletes a driver license from the database based on the license number.
func (repo *SWRepo) DeleteLicense(licenseNumber string) error {
	query := `DELETE FROM driver_licenses WHERE license_number = $1`

	_, err := repo.db.Exec(query, licenseNumber)
	return err
}

// DeletePassport deletes a tech passport from the database based on the passport ID.
func (repo *SWRepo) DeletePassport(passportID string) error {
	query := `DELETE FROM cars WHERE tech_passport_number = $1`

	_, err := repo.db.Exec(query, passportID)
	return err
}
