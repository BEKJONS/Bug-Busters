package models

// DriverLicense represents the driver_licenses table in the database.
type DriverLicense struct {
	ID             string `db:"id"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	FatherName     string `db:"father_name"`
	BirthDate      string `db:"birth_date"` // Changed to string
	Address        string `db:"address"`
	IssueDate      string `db:"issue_date"`      // Changed to string
	ExpirationDate string `db:"expiration_date"` // Changed to string
	Category       string `db:"category"`
	IssuedBy       string `db:"issued_by"`
	LicenseNumber  string `db:"license_number"`
}

type CardId struct {
	ID string `db:"id" json:"id"`
}

type PassportsId struct {
	ID []string `db:"id" json:"id"`
}

type LicenceNumbers struct {
	LicenceNumber string `db:"licence_number" json:"licence_number"`
}

type PassportId struct {
	PassportId string `db:"passport_id" json:"passport_id"`
}
