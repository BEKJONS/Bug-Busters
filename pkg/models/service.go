package models

type Service struct {
	Id                string `json:"id"`
	Type              string `json:"type"`
	Name              string `json:"name"`
	CertificateNumber string `json:"certificate_number"`
	ManagerName       string `json:"manager_name"`
	Address           string `json:"address"`
	PhoneNumber       string `json:"phone_number"`
}

type Services struct {
	Services []Service
}
