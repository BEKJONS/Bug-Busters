package postgres

//import (
//	"bug_busters/pkg/models"
//	"database/sql"
//	"github.com/google/uuid"
//	"testing"
//)
//
//func TestCreateLicense(t *testing.T) {
//	db, err := Connectfortest()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	swRepo := SWRepo{db: db}
//
//	// Test data
//	license := models.DriverLicense{
//		FirstName:      "John",
//		LastName:       "Doe",
//		FatherName:     "DoeFather",
//		BirthDate:      "1990-01-01",
//		Address:        "123 Street",
//		IssueDate:      "2020-01-01",
//		ExpirationDate: "2030-01-01",
//		Category:       "B",
//		IssuedBy:       "DMV",
//		LicenseNumber:  "ABC123456",
//	}
//
//	// Insert driver license
//	err = swRepo.CreateLicense(license)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Verify insertion
//	var count int
//	err = db.Get(&count, `SELECT COUNT(*) FROM driver_licenses WHERE license_number = $1`, license.LicenseNumber)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if count == 0 {
//		t.Errorf("expected license to be created, but it was not found")
//	}
//}
//
//func TestCreatePassport(t *testing.T) {
//	db, err := Connectfortest()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	swRepo := SWRepo{db: db}
//
//	// Insert a car with a known carID
//	carID := uuid.New().String()
//	_, err = db.Exec(`INSERT INTO cars (id) VALUES ($1)`, carID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Call CreatePassport to update the tech passport number
//	err = swRepo.CreatePassport(carID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Verify that the tech_passport_number is updated (it should not be NULL)
//	var techPassportNumber sql.NullString
//	err = db.Get(&techPassportNumber, `SELECT tech_passport_number FROM cars WHERE id = $1`, carID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if !techPassportNumber.Valid {
//		t.Errorf("expected tech_passport_number to be updated, but it is NULL")
//	}
//}
//
//func TestGetLicenseAll(t *testing.T) {
//	db, err := Connectfortest()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	swRepo := SWRepo{db: db}
//
//	// Set up test data
//	licenseNumber := "DEF456789"
//	_, err = db.Exec(`INSERT INTO driver_licenses (license_number, first_name, last_name, father_name, birth_date, address, issue_date, expiration_date, category, issued_by)
//					  VALUES ($1, 'Jane', 'Doe', 'DoeFather', '1985-05-05', '456 Avenue', '2015-05-05', '2025-05-05', 'C', 'DMV')`, licenseNumber)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Test GetLicenseAll
//	licenses, err := swRepo.GetLicenseAll()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if len(licenses) == 0 {
//		t.Errorf("expected at least one license, got none")
//	}
//
//	// Clean up test data
//	_, err = db.Exec(`DELETE FROM driver_licenses WHERE license_number = $1`, licenseNumber)
//	if err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestGetPassportAll(t *testing.T) {
//	db, err := Connectfortest()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	swRepo := SWRepo{db: db}
//
//	// Set up test data
//	passportID := uuid.New().String()
//	_, err = db.Exec(`INSERT INTO cars (tech_passport_number) VALUES ($1)`, passportID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Test GetPassportAll
//	passports, err := swRepo.GetPassportAll()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if len(passports) == 0 {
//		t.Errorf("expected at least one passport, got none")
//	}
//
//	// Clean up test data
//	_, err = db.Exec(`DELETE FROM cars WHERE tech_passport_number = $1`, passportID)
//	if err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDeleteLicense(t *testing.T) {
//	db, err := Connectfortest()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	swRepo := SWRepo{db: db}
//
//	// Set up test data
//	licenseNumber := "GHI789012"
//	_, err = db.Exec(`INSERT INTO driver_licenses (license_number, first_name, last_name, father_name, birth_date, address, issue_date, expiration_date, category, issued_by)
//					  VALUES ($1, 'Bob', 'Smith', 'SmithFather', '1975-03-03', '789 Road', '2005-03-03', '2015-03-03', 'A', 'DMV')`, licenseNumber)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Test DeleteLicense
//	err = swRepo.DeleteLicense(licenseNumber)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Verify deletion
//	var count int
//	err = db.Get(&count, `SELECT COUNT(*) FROM driver_licenses WHERE license_number = $1`, licenseNumber)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if count != 0 {
//		t.Errorf("expected license to be deleted, but it was still found")
//	}
//}
//
//func TestDeletePassport(t *testing.T) {
//	db, err := Connectfortest()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	swRepo := SWRepo{db: db}
//
//	// Set up test data
//	passportID := uuid.New().String()
//	_, err = db.Exec(`INSERT INTO cars (tech_passport_number) VALUES ($1)`, passportID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Test DeletePassport
//	err = swRepo.DeletePassport(passportID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Verify deletion
//	var count int
//	err = db.Get(&count, `SELECT COUNT(*) FROM cars WHERE tech_passport_number = $1`, passportID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if count != 0 {
//		t.Errorf("expected passport to be deleted, but it was still found")
//	}
//}
