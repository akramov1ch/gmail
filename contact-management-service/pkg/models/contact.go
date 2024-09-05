package models

type Contact struct {
	ContactId    string `json:"contact_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserId      string `json:"-"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	IsDeleted   bool   `json:"-"`
	Version     int    `json:"-"`
}