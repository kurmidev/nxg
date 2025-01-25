package repository

import (
	"errors"
	"fmt"
	"log"
	"nxg/internal/domain"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(emp domain.Employee) (domain.Employee, error)
	FindEmployeeById(id uint) (domain.Employee, error)
	UpdateEmployee(id int, emp domain.Employee) (domain.Employee, error)
	GetAllEmployee() ([]domain.Employee, error)
	GetEmployeeByCompany(companyId uint) ([]domain.Employee, error)
	CreateUser(user domain.User) (domain.User, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (e employeeRepository) CreateEmployee(emp domain.Employee) (domain.Employee, error) {
	err := e.db.Create(&emp).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.Employee{}, errors.New("failed to create new employee")
	}
	return emp, nil
}

func (e employeeRepository) FindEmployeeById(id uint) (domain.Employee, error) {
	var emp domain.Employee
	err := e.db.First(&emp, id).Error
	if err != nil {
		log.Printf("find employee by id error %v", err)
		return domain.Employee{}, errors.New("failed to find employee")
	}
	return emp, nil
}

func (e employeeRepository) UpdateEmployee(id int, emp domain.Employee) (domain.Employee, error) {
	err := e.db.Model(&emp).Updates(emp).Error
	if err != nil {
		log.Printf("update employee error %v", err)
		return domain.Employee{}, errors.New("failed to update employee")
	}
	return emp, nil
}

func (e employeeRepository) GetAllEmployee() ([]domain.Employee, error) {
	var ep []domain.Employee
	err := e.db.Find(&ep).Error
	if err != nil {
		log.Printf("get all designations error %v", err)
		return []domain.Employee{}, errors.New("failed to get all product")
	}
	return ep, nil
}

func (e employeeRepository) GetEmployeeByCompany(companyId uint) ([]domain.Employee, error) {
	var ep []domain.Employee
	err := e.db.Find(&ep, "company_id=?", companyId).Error
	if err != nil {
		log.Printf("get all designations error %v", err)
		return []domain.Employee{}, errors.New("failed to get all employee")
	}
	return ep, nil
}

func (d employeeRepository) CreateUser(user domain.User) (domain.User, error) {
	fmt.Printf("user details is %v\n", user)
	err := d.db.Create(&user).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}
	return user, nil
}
