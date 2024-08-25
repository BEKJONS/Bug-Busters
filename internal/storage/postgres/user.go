package postgres

import (
	"bug_busters/pkg/models"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func (u *userRepo) GetProfile(id models.UserId) (models.UserProfile, error) {
	var users models.UserProfile
	var certificate models.DriverLicence

	err := u.db.Get(&users, `SELECT id, driver_license, email, role, created_at, updated_at
									FROM users WHERE id = $1 and deleted_at = 0`, id)
	if err != nil {
		return models.UserProfile{}, err
	}

	err = u.db.Get(&certificate, `SELECT * FROM driver_licenses WHERE license_number = $1`, users.LicenseNumber)
	if err != nil {
		return models.UserProfile{}, err
	}

	return users, nil
}
