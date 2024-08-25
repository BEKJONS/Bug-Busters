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

func (a *AuthRepo) Register(in models.RegisterRequest) (models.RegisterResponse, error) {

	var id string
	query := `INSERT INTO users (phone, email, password,first_name, last_name, username) VALUES ($1, $2, $3,$4, $5, $6) RETURNING id`
	err := a.db.QueryRow(query, in.Phone, in.Email, in.Password, in.FirstName, in.LastName, in.Username).Scan(&id)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return models.RegisterResponse{
		Id:        id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,

		Phone:    in.Phone,
		Username: in.Username,
	}, nil
}
func (a *AuthRepo) LoginEmail(in models.LoginEmailRequest) (models.LoginResponse, error) {

	res := models.LoginResponse{}

	query := `SELECT id, email,username, password,role FROM users WHERE email = $1`
	err := a.db.Get(&res, query, in.Email)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return res, nil
}
