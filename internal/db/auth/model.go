package auth

import (
	"time"
)

type ModelUser struct {
	ID           int        `json:"id"`
	Email        string     `json:"email"`
	Role         string     `json:"role"`
	PasswordHash string     `json:"-"`
	Salt         string     `json:"-"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
