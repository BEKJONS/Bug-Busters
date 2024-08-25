package models

type RegisterRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type LoginEmailRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	Id       string `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type Tokens struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type LicenceNumber struct {
	UserId        string `json:"userId" db:"user_id"`
	LicenceNumber string `json:"number" db:"number"`
}

type Message struct {
	Message string `json:"message" db:"message"`
}
