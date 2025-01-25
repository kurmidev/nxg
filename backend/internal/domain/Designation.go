package domain

import "time"

type Designation struct {
	ID          uint      `json:"id" gorm:"PrimaryKey"`
	Name        string    `json:"name" gorm:"index;unique;not null" `
	DesgFor     int       `json:"desg_for" `
	Description string    `json:"description" `
	ParentId    int       `json:"parent_id"`
	Status      int       `json:"status" gorm:"index"`
	CreatedAt   time.Time `json:"created_at" grorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
}
