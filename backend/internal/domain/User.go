package domain

import "time"

type User struct {
	ID                uint      `json:"id" gorm:"PrimaryKey"`
	Email             string    `json:"email" gorm:"index;unique;not null"`
	UserType          int       `json:"user_type"`
	Password          string    `json:"password"`
	AuthKey           string    `json:"auth_key"`
	PasswordHash      string    `json:"password_hash"`
	VerificationToken string    `json:"verification_token"`
	Status            int       `json:"status" gorm:"default:1"`
	LastLoginAt       time.Time `json:"last_login_at" gorm:"default:null"`
	CreatedAt         time.Time `json:"created_at" grorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy         int       `json:"created_by"`
	UpdatedBy         int       `json:"updated_by"`
}
