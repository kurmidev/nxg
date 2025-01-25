package repository

import (
	"errors"
	"log"
	"nxg/internal/domain"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	CreateCompany(c domain.Company) (domain.Company, error)
	FindCompanyById(id int) (domain.Company, error)
	UpdateCompany(id int, d domain.Company) (domain.Company, error)
	GetAllCompany() ([]domain.Company, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (c companyRepository) CreateCompany(comp domain.Company) (domain.Company, error) {
	err := c.db.Create(&comp).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.Company{}, errors.New("failed to create new designation")
	}
	return comp, nil
}

func (c companyRepository) FindCompanyById(id int) (domain.Company, error) {
	var comp domain.Company
	err := c.db.First(&comp, id).Error
	if err != nil {
		log.Printf("find company by id error %v", err)
		return domain.Company{}, errors.New("failed to find company")
	}
	return comp, nil
}

func (c companyRepository) UpdateCompany(id int, comp domain.Company) (domain.Company, error) {
	err := c.db.Model(&comp).Updates(comp).Error
	if err != nil {
		log.Printf("update company error %v", err)
		return domain.Company{}, errors.New("failed to update company details")
	}
	return comp, nil
}

func (c companyRepository) GetAllCompany() ([]domain.Company, error) {
	var comp []domain.Company
	err := c.db.Find(&comp).Error
	if err != nil {
		log.Printf("get all designations error %v", err)
		return []domain.Company{}, errors.New("failed to get all company")
	}
	return comp, nil
}
