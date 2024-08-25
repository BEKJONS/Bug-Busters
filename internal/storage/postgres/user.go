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

func (u *userRepo) AddImage(in *models.UpdateCarImage) error {

	_, err := u.db.Exec("update cars set image_url = $1 where user_id = $2", in.Url, in.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) GetImage(useId string) (string, error) {
	var url string
	err := u.db.Get(&url, `SELECT image_url FROM cars WHERE user_id = $1`, useId)
	if err != nil {
		return "", err
	}

	return url, nil
}
