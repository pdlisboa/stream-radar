package model

import "time"

type User struct {
	Id           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"not null" json:"email"`
	Name         string    `gorm:"not null" json:"name"`
	PasswordHash string    `gorm:"not null" json:"password_hash,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
