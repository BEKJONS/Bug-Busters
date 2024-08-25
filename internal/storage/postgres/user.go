package postgres

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) storage.UserStorage {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetProfile(id models.UserId) (models.UserProfile, error) {
	var users models.UserProfile
	var certificate models.DriverLicence

	err := u.db.Get(&users, `SELECT id, driver_license, email, role, created_at, updated_at
									FROM users WHERE id = $1 and deleted_at = 0`, id.Id)
	if err != nil {
		return models.UserProfile{}, err
	}

	err = u.db.Get(&certificate, `SELECT * FROM driver_licenses WHERE license_number = $1`, users.DriverLicense)
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

func (u *userRepo) GetPaidFines(userId string) (*[]*models.UserFines, error) {
	res := []*models.UserFines{}

	err := u.db.Select(&res, "SELECT officer_id, license_plate FROM fines WHERE fine_owner = $1 AND payment_date IS NOT NULL", userId)
	if err != nil {
		return nil, err
	}

	var userName string
	err = u.db.Get(&userName, `SELECT d.first_name FROM users u INNER JOIN driver_licenses d
	                           ON u.driver_license = d.license_number WHERE u.id = $1`, userId)
	if err != nil {
		return nil, err
	}

	for _, fine := range res {
		fine.CarOwnerName = userName
	}

	return &res, nil
}

func (u *userRepo) GetUnpaid(userId string) (*[]*models.UserFines, error) {
	res := []*models.UserFines{}

	err := u.db.Select(&res, "SELECT officer_id, license_plate FROM fines WHERE fine_owner = $1 AND payment_date IS NULL", userId)
	if err != nil {
		return nil, err
	}

	var userName string
	err = u.db.Get(&userName, `SELECT d.first_name FROM users u INNER JOIN driver_licenses d
	                           ON u.driver_license = d.license_number WHERE u.id = $1`, userId)
	if err != nil {
		return nil, err
	}

	for _, fine := range res {
		fine.CarOwnerName = userName
	}

	return &res, nil
}

func (u *userRepo) DeleteUser(userId string) error {
	_, err := u.db.Exec(`UPDATE users 
	                      SET deleted_at = date_part('epoch', current_timestamp)::INT 
	                      WHERE id = $1 AND deleted_at = 0`, userId)
	return err
}
