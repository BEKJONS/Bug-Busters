package models

type Fine struct {
	ID                 string `json:"id" db:"id"`
	TechPassportNumber string `json:"tech_passport_number" db:"tech_passport_number"`
	LicensePlate       string `json:"license_plate" db:"license_plate"`
	OfficerID          string `json:"officer_id" db:"officer_id"`
	FineOwner          string `json:"fine_owner" db:"fine_owner"`
	FineDate           string `json:"fine_date" db:"fine_date"`
	PaymentDate        string `json:"payment_date" db:"payment_date"`
	Price              int    `json:"price" db:"price"`
}
type FineReq struct {
	TechPassportNumber string `json:"tech_passport_number" db:"tech_passport_number"`
	LicensePlate       string `json:"license_plate" db:"license_plate"`
	OfficerID          string `json:"officer_id" db:"officer_id"`
	FineOwner          string `json:"fine_owner" db:"fine_owner"`
	FineDate           string `json:"fine_date" db:"fine_date"`
	PaymentDate        string `json:"payment_date" db:"payment_date"`
	Price              int    `json:"price" db:"price"`
}
type FineAccept struct {
	Id          string `json:"id" db:"id"`
	PaymentDate string `json:"payment_date" db:"payment_date"`
}
type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
type Fines struct {
	Fine Fine `json:"fine"`
}
