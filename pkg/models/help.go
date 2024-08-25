package models

type Message struct {
	Message string `json:"message" db:"message"`
}

type Error struct {
	Error string `json:"error" db:"error"`
}
