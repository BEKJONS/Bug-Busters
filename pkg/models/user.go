package models

type UserId struct {
	Id string `json:"id"`
}

type UserProfile struct {
	Id            string        `json:"user_id" db:"id"`
	DriverLicense string        `json:"driver_license" db:"driver_license"`
	Email         string        `json:"email" db:"email"`
	Role          string        `json:"role" db:"role"`
	CreatedAt     string        `json:"created_at" db:"created_at"`
	UpdatedAt     string        `json:"updated_at" db:"updated_at"`
	Driver        DriverLicence `json:"driver_license"`
}

type DriverLicence struct {
	Id             string `json:"id" db:"id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	FatherName     string `json:"father_name" db:"father_name"`
	BirthDate      string `json:"birth_date" db:"birth_date"`
	Address        string `json:"address" db:"address"`
	IssueDate      string `json:"issue_date" db:"issue_date"`
	ExpirationDate string `json:"expiration_date" db:"expiration_date"`
	Category       string `json:"category" db:"category"`
	IssuedBy       string `json:"issued_by" db:"issued_by"`
	LicenceNumber  string `json:"licence_number" db:"license_number"`
}

type UpdateCarImage struct {
	UserId string `json:"user_id" db:"user_id"`
	Url    string `json:"url" db:"url"`
}

type Url struct {
	Url string `json:"url" db:"url"`
}
