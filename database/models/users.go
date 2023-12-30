package models

import "github.com/google/uuid"

type User struct {
	ID       *uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid();"`
	Username *string    `gorm:"unique;not null" json:"username"`
	Email    *string    `gorm:"unique;not null" json:"email"`
	Phone    *string    `gorm:"unique;not null" json:"phone"`
}
