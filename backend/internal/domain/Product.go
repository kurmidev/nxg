package domain

import "time"

type Product struct {
	ID   uint   `json:"id" gorm:"PrimaryKey"`
	Name string `json:"name" gorm:"index;unique;not null"`
	//ProjectAtrribute []ProductAttribute `json:"project_attribute"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at" grorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
}

type ProductAttribute struct {
	ID        uint      `json:"id" gorm:"PrimaryKey"`
	Name      string    `json:"name" gorm:"index"`
	Values    string    `json:"values" gorm:"not null"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at" grorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy int       `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
}
