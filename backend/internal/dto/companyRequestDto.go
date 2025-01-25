package dto

type CompanyCreateDto struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"  `
	Address       string `json:"address" `
	Email         string `json:"email"`
	PhoneNo       string `json:"phone"`
	MobileNo      string `json:"mobile"`
	ContactPerson string `json:"contact"`
}

type CompanyUpdateDto struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"  `
	Address       string `json:"address" `
	Email         string `json:"email" `
	PhoneNo       string `json:"phone"`
	MobileNo      string `json:"mobile"`
	ContactPerson string `json:"contact"`
	Status        int    `json:"status"`
}
