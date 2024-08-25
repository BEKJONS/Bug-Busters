package postgres

import (
	"bug_busters/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

// ConnectDB sets up a connection to the PostgreSQL database
func ConnectDB() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "BEKJONS", "road_24")

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// TestCreateFines tests the creation of a fine
func TestCreateFines(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewIIRepo(db)

	fine := &models.FineReq{
		TechPassportNumber: "a9efb02b-e778-41c4-a9c6-4d92e16dd0d1",
		LicensePlate:       "ABC123",
		OfficerID:          "0de5b099-5ca1-43e6-9659-03bee85901cb",
		FineOwner:          "2f0f677f-4db7-49cb-93d5-01a43521208c",
		Price:              150,
	}

	err = repo.CreateFines(fine)
	if err != nil {
		t.Fatalf("Failed to create fine: %v", err)
	}

	fmt.Println("Fine created successfully")
}

// TestAcceptFinesById tests accepting a fine by ID
func TestAcceptFinesById(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewIIRepo(db)

	accept := models.FineAccept{
		Id: "04739238-a219-4b5a-8ccf-e156e8e5b406",
	}

	err = repo.AcceptFinesById(accept)
	if err != nil {
		t.Fatalf("Failed to accept fine: %v", err)
	}

	fmt.Println("Fine accepted successfully")
}

// TestGetPaidFines tests retrieving paid fines
func TestGetPaidFines(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewIIRepo(db)

	pagination := models.Pagination{
		Page:  1,
		Limit: 10,
	}

	fines, err := repo.GetPaidFines(pagination)
	if err != nil {
		t.Fatalf("Failed to get paid fines: %v", err)
	}

	fmt.Printf("Paid Fines: %+v\n", fines)
}

// TestGetUnpaidFines tests retrieving unpaid fines
func TestGetUnpaidFines(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewIIRepo(db)

	pagination := models.Pagination{
		Page:  1,
		Limit: 10,
	}

	fines, err := repo.GetUnpaidFines(pagination)
	if err != nil {
		t.Fatalf("Failed to get unpaid fines: %v", err)
	}

	fmt.Printf("Unpaid Fines: %+v\n", fines)
}

// TestGetAllFines tests retrieving all fines
func TestGetAllFines(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewIIRepo(db)

	pagination := models.Pagination{
		Page:  1,
		Limit: 10,
	}

	fines, err := repo.GetAllFines(pagination)
	if err != nil {
		t.Fatalf("Failed to get all fines: %v", err)
	}

	fmt.Printf("All Fines: %+v\n", fines)
}
