package postgres

import (
	"bug_busters/internal/storage"
	"bug_busters/pkg/models"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type ServiceRepo struct {
	db *sqlx.DB
}

func NewServiceRepo(db *sqlx.DB) storage.ServiceStorage {
	return &ServiceRepo{
		db: db,
	}
}

func (r *ServiceRepo) CreateService(service *models.Service) (*models.Service, error) {
	query := `INSERT INTO Services (type, name, certificate_number, manager_name, address, phone_number) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRow(query, service.Type, service.Name, service.CertificateNumber, service.ManagerName, service.Address, service.PhoneNumber).Scan(&service.Id)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (r *ServiceRepo) GetService(id string) (*models.Service, error) {
	service := &models.Service{}
	query := `SELECT id, type, name, certificate_number, manager_name, address, phone_number FROM Services WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&service.Id, &service.Type, &service.Name, &service.CertificateNumber, &service.ManagerName, &service.Address, &service.PhoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return service, nil
}

func (r *ServiceRepo) UpdateService(service *models.Service) (*models.Service, error) {
	query := `UPDATE Services SET type = $1, name = $2, certificate_number = $3, manager_name = $4, address = $5, phone_number = $6 WHERE id = $7`
	_, err := r.db.Exec(query, service.Type, service.Name, service.CertificateNumber, service.ManagerName, service.Address, service.PhoneNumber, service.Id)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (r *ServiceRepo) DeleteService(id string) (string, error) {
	query := `DELETE FROM Services WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return "failed to delete", err
	}
	return "Service deleted", nil
}

func (r *ServiceRepo) GetServices() (*models.Services, error) {
	var services []models.Service
	query := `SELECT id, type, name, certificate_number, manager_name, address, phone_number FROM Services`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.Id, &service.Type, &service.Name, &service.CertificateNumber, &service.ManagerName, &service.Address, &service.PhoneNumber)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return &models.Services{Services: services}, nil
}
