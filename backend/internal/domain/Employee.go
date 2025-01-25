package domain

import "time"

type Employee struct {
	ID            uint      `json:"id" gorm:"PrimaryKey"`
	Name          string    `json:"name"`
	Email         string    `json:"email" gorm:"index;unique;not null"`
	MobileNo      string    `json:"mobile_no" gorm:"index;unique;not null"`
	DesignationId uint      `json:"designation_id"`
	Role          int       `json:"role"`
	CompanyId     uint      `json:"company_id"`
	Status        int       `json:"status" gorm:"default:1"`
	CreatedAt     time.Time `json:"created_at" grorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy     int       `json:"created_by"`
	UpdatedBy     int       `json:"updated_by"`
}
