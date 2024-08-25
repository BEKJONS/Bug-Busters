package postgres

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) storage.AuthStorage {
	return &AuthRepo{
		db: db,
	}
}

func (a *AuthRepo) Register(in models.RegisterRequest) error {

	query := `INSERT INTO users (email,password,role) VALUES ($1, $2, $3)`
	err := a.db.QueryRow(query, in.Email, in.Password, in.Role)
	if err != nil {
		return err.Err()
	}

	return nil
}
func (a *AuthRepo) Login(in *models.LoginEmailRequest) (*models.LoginResponse, error) {

	res := &models.LoginResponse{}

	query := `SELECT id, password, role FROM users WHERE email = $1`
	err := a.db.Get(res, query, in.Email)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AuthRepo) AddLicence(in *models.LicenceNumber) error {
	_, err := a.db.Exec("UPDATE users set driver_license = $1 where id = $2 and deleted_at = 0", in.LicenceNumber, in.UserId)
	if err != nil {
		return err
	}

	return nil
}
