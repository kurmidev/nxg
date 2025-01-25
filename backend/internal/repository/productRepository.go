package repository

import (
	"errors"
	"log"
	"nxg/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(c domain.Product) (domain.Product, error)
	FindProductById(id int) (domain.Product, error)
	UpdateProduct(id int, d domain.Product) (domain.Product, error)
	GetAllProduct() ([]domain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p productRepository) CreateProduct(comp domain.Product) (domain.Product, error) {
	err := p.db.Create(&comp).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.Product{}, errors.New("failed to create new product")
	}
	return comp, nil
}

func (p productRepository) FindProductById(id int) (domain.Product, error) {
	var prod domain.Product
	err := p.db.First(&prod, id).Error
	if err != nil {
		log.Printf("find company by id error %v", err)
		return domain.Product{}, errors.New("failed to find product")
	}
	return prod, nil
}

func (p productRepository) UpdateProduct(id int, prod domain.Product) (domain.Product, error) {
	err := p.db.Model(&prod).Updates(prod).Error
	if err != nil {
		log.Printf("update product error %v", err)
		return domain.Product{}, errors.New("failed to update product")
	}
	return prod, nil
}

func (p productRepository) GetAllProduct() ([]domain.Product, error) {
	var prod []domain.Product
	err := p.db.Find(&prod).Error
	if err != nil {
		log.Printf("get all product error %v", err)
		return []domain.Product{}, errors.New("failed to get all product")
	}
	return prod, nil
}
