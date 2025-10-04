package model

import (
	"stream-radar/domain/dto"
	"time"
)

type User struct {
	Id           uint       `gorm:"primaryKey" json:"id"`
	Email        string     `gorm:"not null" json:"email"`
	Name         string     `gorm:"not null" json:"name"`
	PasswordHash string     `gorm:"not null" json:"password_hash,omitempty"`
	Streamers    []Streamer `gorm:"many2many:user_streamers;" json:"streamers"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (user User) ToDTO() dto.UserDTO {
	return dto.UserDTO{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (User) FromDto(user dto.UserDTO) User {
	return User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
