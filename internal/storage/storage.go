package storage

import "bug_busters/pkg/models"

type AuthStorage interface {
	Register(in models.RegisterRequest) error
	Login(in *models.LoginEmailRequest) (*models.LoginResponse, error)
	AddLicence(in *models.LicenceNumber) error
}

type IIStorage interface {
	CreateFines(fine *models.FineReq) error
	AcceptFinesById(accept models.FineAccept) error
	GetPaidFines(pagination models.Pagination) (models.Fines, error)
	GetUnpaidFines(pagination models.Pagination) (models.Fines, error)
	GetAllFines(pagination models.Pagination) (*models.Fines, error)
}

type UserStorage interface {
	GetProfile(id models.UserId) (models.UserProfile, error)
	AddImage(in *models.UpdateCarImage) error
	GetImage(useId string) (string, error)
	GetPaidFinesU(userId string) (*[]*models.UserFines, error)
	GetUnpaid(userId string) (*[]*models.UserFines, error)
	DeleteUser(userId string) error
}

type SWStorage interface {
	CreateLicense(license models.DriverLicense) error
	CreatePassport(CarId string) error
	GetLicenseAll() ([]models.DriverLicense, error)
	GetPassportAll() ([]string, error)
	DeleteLicense(licenseNumber string) error
	DeletePassport(passportID string) error
}
