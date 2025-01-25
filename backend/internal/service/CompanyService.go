package service

import (
	"errors"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/helper"
	"nxg/internal/repository"
	"time"
)

type CompanyService struct {
	Repo         repository.CompanyRepository
	Auth         helper.Auth
	LoggedInUser domain.User
}

func (c CompanyService) GetCompany() ([]domain.Company, error) {
	company, err := c.Repo.GetAllCompany()
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c CompanyService) GetCompanyById(id int) (*domain.Company, error) {
	company, err := c.Repo.FindCompanyById(id)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (c CompanyService) CreateCompany(comp *dto.CompanyCreateDto) (*domain.Company, error) {
	company := &domain.Company{
		Name:          comp.Name,
		Address:       comp.Address,
		Email:         comp.Email,
		PhoneNo:       comp.PhoneNo,
		MobileNo:      comp.MobileNo,
		ContactPerson: comp.ContactPerson,
		CreatedAt:     time.Now(),
		CreatedBy:     int(c.LoggedInUser.ID),
	}
	newcompany, err := c.Repo.CreateCompany(*company)
	if err != nil {
		return nil, err
	}
	return &newcompany, nil
}

func (c CompanyService) UpdateCompany(id int, comp *dto.CompanyUpdateDto) (*domain.Company, error) {
	oldComp, err := c.Repo.FindCompanyById(id)
	if err != nil {
		return nil, errors.New("no designation found")
	}

	ucomp, err := c.Repo.UpdateCompany(id, domain.Company{
		ID:            uint(id),
		Name:          comp.Name,
		Address:       comp.Address,
		Email:         comp.Email,
		PhoneNo:       comp.PhoneNo,
		MobileNo:      comp.MobileNo,
		ContactPerson: comp.ContactPerson,
		UpdatedBy:     int(c.LoggedInUser.ID),
		UpdatedAt:     time.Now(),
		Status:        comp.Status,
		CreatedBy:     oldComp.CreatedBy,
		CreatedAt:     oldComp.CreatedAt,
	})
	if err != nil {
		return nil, errors.New("failed to update company details")
	}

	return &ucomp, nil
}
