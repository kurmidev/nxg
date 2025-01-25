package service

import (
	"errors"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/helper"
	"nxg/internal/repository"
	"time"
)

type ProductService struct {
	Repo         repository.ProductRepository
	Auth         helper.Auth
	LoggedInUser domain.User
}

func (d ProductService) GetProducts() ([]domain.Product, error) {
	Products, err := d.Repo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (d ProductService) GetProductById(id int) (*domain.Product, error) {
	Product, err := d.Repo.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return &Product, nil
}

func (p ProductService) CreateProduct(prod *dto.ProductCreateDto) (*domain.Product, error) {
	product := &domain.Product{
		Name:        prod.Name,
		Description: prod.Description,
		CreatedAt:   time.Now(),
		CreatedBy:   int(p.LoggedInUser.ID),
	}
	newproduct, err := p.Repo.CreateProduct(*product)
	if err != nil {
		return nil, err
	}
	return &newproduct, nil
}

func (p ProductService) UpdateProduct(id int, prod *dto.ProductUpdateDto) (*domain.Product, error) {
	oldProd, err := p.Repo.FindProductById(id)
	if err != nil {
		return nil, errors.New("no product found")
	}

	uproduct, err := p.Repo.UpdateProduct(id, domain.Product{
		ID:          uint(id),
		Name:        prod.Name,
		Description: prod.Description,
		UpdatedBy:   int(p.LoggedInUser.ID),
		UpdatedAt:   time.Now(),
		CreatedAt:   oldProd.CreatedAt,
		CreatedBy:   oldProd.CreatedBy,
	})
	if err != nil {
		return nil, errors.New("failed to update product details")
	}

	return &uproduct, nil
}
