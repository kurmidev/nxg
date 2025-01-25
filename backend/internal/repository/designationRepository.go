package repository

import (
	"errors"
	"log"
	"nxg/internal/domain"

	"gorm.io/gorm"
)

type DesignationRepository interface {
	CreateDesignation(d domain.Designation) (domain.Designation, error)
	FindDesignationById(id int) (domain.Designation, error)
	UpdateDesignation(id int, d domain.Designation) (domain.Designation, error)
	GetAllDesignations() ([]domain.Designation, error)
}

type designationRepository struct {
	db *gorm.DB
}

func NewDesignationRepository(db *gorm.DB) DesignationRepository {
	return &designationRepository{db: db}
}

func (d designationRepository) CreateDesignation(desg domain.Designation) (domain.Designation, error) {
	err := d.db.Create(&desg).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.Designation{}, errors.New("failed to create new designation")
	}
	return desg, nil
}

func (d designationRepository) FindDesignationById(id int) (domain.Designation, error) {
	var desg domain.Designation
	err := d.db.First(&desg, id).Error
	if err != nil {
		log.Printf("find user by id error %v", err)
		return domain.Designation{}, errors.New("failed to find designation")
	}
	return desg, nil
}

func (d designationRepository) UpdateDesignation(id int, desg domain.Designation) (domain.Designation, error) {
	//err := d.db.Model(&des).Clauses(clause.Returning{}).Where("id=?", id).Updates(desg).Error
	//err := d.db.Save(&desg).Error
	err := d.db.Model(&desg).Updates(desg).Error
	if err != nil {
		log.Printf("update designation error %v \n %v", err, desg)
		return domain.Designation{}, errors.New("failed to update user")
	}
	return desg, nil
}

func (d designationRepository) GetAllDesignations() ([]domain.Designation, error) {
	var desgs []domain.Designation
	err := d.db.Find(&desgs).Error
	if err != nil {
		log.Printf("get all designations error %v", err)
		return []domain.Designation{}, errors.New("failed to get all designations")
	}
	return desgs, nil
}
