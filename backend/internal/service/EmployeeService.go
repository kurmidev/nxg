package service

import (
	"errors"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/helper"
	"nxg/internal/repository"
	"time"
)

type EmployeeService struct {
	Repo         repository.EmployeeRepository
	Auth         helper.Auth
	LoggedInUser domain.User
}

func (e EmployeeService) GetEmployees() ([]domain.Employee, error) {
	Employees, err := e.Repo.GetAllEmployee()
	if err != nil {
		return nil, err
	}
	return Employees, nil
}

func (e EmployeeService) GetEmployeeById(id int) (*domain.Employee, error) {
	Employee, err := e.Repo.FindEmployeeById(uint(id))
	if err != nil {
		return nil, err
	}
	return &Employee, nil
}

func (e EmployeeService) CreateEmployee(Employee *dto.EmployeeCreateDto) (*domain.Employee, error) {
	emp := domain.Employee{
		Name:          Employee.Name,
		Email:         Employee.Email,
		MobileNo:      Employee.MobileNo,
		DesignationId: Employee.DesignationId,
		Role:          Employee.Role,
		CompanyId:     Employee.CompanyId,
		CreatedAt:     time.Now(),
		CreatedBy:     int(e.LoggedInUser.ID),
	}

	employee, err := e.Repo.CreateEmployee(emp)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e EmployeeService) UpdateEmployee(id int, emp *dto.EmployeeUpdateDto) (*domain.Employee, error) {
	oldEmp, err := e.Repo.FindEmployeeById(uint(id))
	if err != nil {
		return nil, errors.New("no employee found")
	}

	uemp := domain.Employee{
		ID:            uint(id),
		Name:          emp.Name,
		Email:         emp.Email,
		MobileNo:      emp.MobileNo,
		DesignationId: emp.DesignationId,
		Role:          emp.Role,
		CompanyId:     emp.CompanyId,
		Status:        1,
		UpdatedBy:     int(e.LoggedInUser.ID),
		UpdatedAt:     time.Now(),
		CreatedAt:     oldEmp.CreatedAt,
		CreatedBy:     oldEmp.CreatedBy,
	}

	employee, err := e.Repo.UpdateEmployee(id, uemp)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
