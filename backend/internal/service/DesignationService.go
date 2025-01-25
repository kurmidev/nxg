package service

import (
	"errors"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/helper"
	"nxg/internal/repository"
	"time"
)

type DesignationService struct {
	Repo         repository.DesignationRepository
	Auth         helper.Auth
	LoggedInUser domain.User
}

func (d DesignationService) GetDesignations() ([]domain.Designation, error) {
	designations, err := d.Repo.GetAllDesignations()
	if err != nil {
		return nil, err
	}
	return designations, nil
}

func (d DesignationService) GetDesignationById(id int) (*domain.Designation, error) {
	designation, err := d.Repo.FindDesignationById(id)
	if err != nil {
		return nil, err
	}
	return &designation, nil
}

func (d DesignationService) CreateDesignation(designation *dto.DesignationCreateDto) (*domain.Designation, error) {
	desg := domain.Designation{
		Name:        designation.Name,
		DesgFor:     designation.DesgFor,
		Description: designation.Description,
		ParentId:    designation.ParentId,
		CreatedAt:   time.Now(),
		CreatedBy:   int(d.LoggedInUser.ID),
	}
	des, err := d.Repo.CreateDesignation(desg)
	if err != nil {
		return nil, err
	}
	return &des, nil
}

func (d DesignationService) UpdateDesignation(id int, designation *dto.DesignationUpdateDto) (*domain.Designation, error) {
	olddesg, err := d.Repo.FindDesignationById(id)
	if err != nil {
		return nil, errors.New("no designation found")
	}

	newdeg := domain.Designation{
		ID:          uint(id),
		Name:        designation.Name,
		DesgFor:     designation.DesgFor,
		Description: designation.Description,
		ParentId:    designation.ParentId,
		UpdatedBy:   int(d.LoggedInUser.ID),
		UpdatedAt:   time.Now(),
		Status:      designation.Status, // Assuming status is active by default for now. Will update this if required.
		CreatedBy:   olddesg.CreatedBy,  // Assuming created_by is not changed. Will update this if required.
		CreatedAt:   olddesg.CreatedAt,
	}

	udesg, err := d.Repo.UpdateDesignation(id, newdeg)
	if err != nil {
		return nil, errors.New("failed to update designation")
	}

	return &udesg, nil
}
