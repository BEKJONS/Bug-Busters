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
