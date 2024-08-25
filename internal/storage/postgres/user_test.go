package postgres

import (
	"bug_busters/pkg/models"
	"github.com/google/uuid"
	"testing"
)

func TestGetProfile(t *testing.T) {
	db, err := Connectfortest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepo := userRepo{db: db}

	// Set up test data
	licenseNumber := "ABC123456"
	_, err = db.Exec(`INSERT INTO driver_licenses (license_number, first_name, last_name, father_name, birth_date, address, issue_date, expiration_date, category, issued_by) 
						VALUES ($1, 'John', 'Doe', 'DoeFather', '1990-01-01', '123 Street', '2020-01-01', '2030-01-01', 'B', 'DMV')`, licenseNumber)
	if err != nil {
		t.Fatal(err)
	}

	userID := uuid.New().String()
	_, err = db.Exec(`INSERT INTO users (id, email, driver_license, role, created_at, updated_at, deleted_at) 
						VALUES ($1, 'test@example.com', $2, 'user', NOW(), NOW(), 0)`, userID, licenseNumber)
	if err != nil {
		t.Fatal(err)
	}

	id := models.UserId{Id: userID}
	profile, err := userRepo.GetProfile(id)
	if err != nil {
		t.Fatal(err)
	}

	if profile.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got %s", profile.Email)
	}

	// Clean up test data
	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM driver_licenses WHERE license_number = $1`, licenseNumber)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddImage(t *testing.T) {
	db, err := Connectfortest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepo := userRepo{db: db}

	// Set up test data
	userID := uuid.New().String()
	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at) VALUES ($1, 'carowner@example.com', 'owner', NOW(), NOW(), 0)`, userID)
	if err != nil {
		t.Fatal(err)
	}

	carID := uuid.New().String()
	_, err = db.Exec(`INSERT INTO cars (id, user_id, image_url) VALUES ($1, $2, '')`, carID, userID)
	if err != nil {
		t.Fatal(err)
	}

	// Test AddImage
	updateImage := &models.UpdateCarImage{
		Url:    "http://example.com/car.jpg",
		UserId: userID,
	}
	err = userRepo.AddImage(updateImage)
	if err != nil {
		t.Fatal(err)
	}

	var imageURL string
	err = db.Get(&imageURL, `SELECT image_url FROM cars WHERE user_id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}

	if imageURL != "http://example.com/car.jpg" {
		t.Errorf("expected image_url 'http://example.com/car.jpg', got %s", imageURL)
	}

	// Clean up test data
	_, err = db.Exec(`DELETE FROM cars WHERE user_id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetImage(t *testing.T) {
	db, err := Connectfortest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepo := userRepo{db: db}

	// Set up test data
	userID := uuid.New().String()
	imageURL := "http://example.com/car.jpg"
	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at) VALUES ($1, 'carowner@example.com', 'owner', NOW(), NOW(), 0)`, userID)
	if err != nil {
		t.Fatal(err)
	}

	carID := uuid.New().String()
	_, err = db.Exec(`INSERT INTO cars (id, user_id, image_url) VALUES ($1, $2, $3)`, carID, userID, imageURL)
	if err != nil {
		t.Fatal(err)
	}

	// Test GetImage
	result, err := userRepo.GetImage(userID)
	if err != nil {
		t.Fatal(err)
	}

	if result != imageURL {
		t.Errorf("expected image_url '%s', got %s", imageURL, result)
	}

	// Clean up test data
	_, err = db.Exec(`DELETE FROM cars WHERE user_id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPaidFinesU(t *testing.T) {
	db, err := Connectfortest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepo := userRepo{db: db}

	// Set up test data
	userID := uuid.New().String()
	officerID := uuid.New().String()
	licensePlate := "ABC123"
	driverLicenseNumber := "DL12345"

	// Insert driver license for user
	_, err = db.Exec(`INSERT INTO driver_licenses (license_number, first_name, last_name, issue_date, expiration_date) VALUES ($1, 'John', 'Doe', NOW(), NOW() + interval '1 year')`, driverLicenseNumber)
	if err != nil {
		t.Fatal(err)
	}

	// Insert users
	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at, driver_license) VALUES ($1, 'fineowner@example.com', 'owner', NOW(), NOW(), 0, $2)`, userID, driverLicenseNumber)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at) VALUES ($1, 'officer@example.com', 'officer', NOW(), NOW(), 0)`, officerID)
	if err != nil {
		t.Fatal(err)
	}

	// Insert car for user
	_, err = db.Exec(`INSERT INTO cars (id, user_id, license_plate) VALUES (gen_random_uuid(), $1, $2)`, userID, licensePlate)
	if err != nil {
		t.Fatal(err)
	}

	// Insert paid fine
	_, err = db.Exec(`INSERT INTO fines (id, fine_owner, officer_id, license_plate, fine_date, price, payment_date) VALUES (gen_random_uuid(), $1, $2, $3, NOW(), 100, NOW())`, userID, officerID, licensePlate)
	if err != nil {
		t.Fatal(err)
	}

	// Test GetPaidFines
	fines, err := userRepo.GetPaidFines(userID)
	if err != nil {
		t.Fatal(err)
	}

	if len(*fines) == 0 {
		t.Errorf("expected at least one paid fine, got none")
	}

	// Check that the data is correct
	for _, fine := range *fines {
		if fine.OfficerId != officerID || fine.LicencePlate != licensePlate {
			t.Errorf("unexpected fine data: %+v", fine)
		}
	}

	// Clean up test data
	_, err = db.Exec(`DELETE FROM fines WHERE fine_owner = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM cars WHERE user_id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM users WHERE id IN ($1, $2)`, userID, officerID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM driver_licenses WHERE license_number = $1`, driverLicenseNumber)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUnpaited(t *testing.T) {
	db, err := Connectfortest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepo := userRepo{db: db}

	// Set up test data
	userID := uuid.New().String()
	officerID := uuid.New().String()
	licensePlate := "ABC123"
	driverLicenseNumber := "DL12345"

	// Insert driver license for user
	_, err = db.Exec(`INSERT INTO driver_licenses (license_number, first_name, last_name, issue_date, expiration_date) VALUES ($1, 'John', 'Doe', NOW(), NOW() + interval '1 year')`, driverLicenseNumber)
	if err != nil {
		t.Fatal(err)
	}

	// Insert users
	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at, driver_license) VALUES ($1, 'fineowner@example.com', 'owner', NOW(), NOW(), 0, $2)`, userID, driverLicenseNumber)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at) VALUES ($1, 'officer@example.com', 'officer', NOW(), NOW(), 0)`, officerID)
	if err != nil {
		t.Fatal(err)
	}

	// Insert car for user
	_, err = db.Exec(`INSERT INTO cars (id, user_id, license_plate) VALUES (gen_random_uuid(), $1, $2)`, userID, licensePlate)
	if err != nil {
		t.Fatal(err)
	}

	// Insert unpaid fine
	_, err = db.Exec(`INSERT INTO fines (id, fine_owner, officer_id, license_plate, fine_date, price, payment_date) VALUES (gen_random_uuid(), $1, $2, $3, NOW(), 100, NULL)`, userID, officerID, licensePlate)
	if err != nil {
		t.Fatal(err)
	}

	// Test GetUnpaited
	fines, err := userRepo.GetUnpaited(userID)
	if err != nil {
		t.Fatal(err)
	}

	if len(*fines) == 0 {
		t.Errorf("expected at least one unpaid fine, got none")
	}

	// Check that the data is correct
	for _, fine := range *fines {
		if fine.OfficerId != officerID || fine.LicencePlate != licensePlate {
			t.Errorf("unexpected fine data: %+v", fine)
		}
		if fine.CarOwnerName != "John" {
			t.Errorf("expected car owner name to be 'John', got '%s'", fine.CarOwnerName)
		}
	}

	// Clean up test data
	_, err = db.Exec(`DELETE FROM fines WHERE fine_owner = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM cars WHERE user_id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM users WHERE id IN ($1, $2)`, userID, officerID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM driver_licenses WHERE license_number = $1`, driverLicenseNumber)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := Connectfortest()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	userRepo := userRepo{db: db}

	// Set up test data
	userID := uuid.New().String()

	// Insert a user with deleted_at set to 0
	_, err = db.Exec(`INSERT INTO users (id, email, role, created_at, updated_at, deleted_at) VALUES ($1, 'test@example.com', 'test_role', NOW(), NOW(), 0)`, userID)
	if err != nil {
		t.Fatal(err)
	}

	// Call DeleteUser
	err = userRepo.DeleteUser(userID)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the user has been deleted (deleted_at should not be 0)
	var deletedAt int64
	err = db.Get(&deletedAt, `SELECT deleted_at FROM users WHERE id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}

	if deletedAt == 0 {
		t.Errorf("expected deleted_at to be updated, got 0")
	}

	// Clean up test data if necessary
	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, userID)
	if err != nil {
		t.Fatal(err)
	}
}
