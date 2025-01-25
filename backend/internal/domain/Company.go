package domain

import "time"

type Company struct {
	ID            uint      `json:"id" gorm:"PrimaryKey"`
	Name          string    `json:"name"  `
	Address       string    `json:"address" `
	Email         string    `json:"email" gorm:"index;unique;not null"`
	PhoneNo       string    `json:"phone"`
	MobileNo      string    `json:"mobile"`
	ContactPerson string    `json:"contact"`
	Status        int       `json:"status" gorm:"index"`
	CreatedAt     time.Time `json:"created_at" grorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy     int       `json:"created_by"`
	UpdatedBy     int       `json:"updated_by"`
}
