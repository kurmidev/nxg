package handlers

import (
	"errors"
	"net/http"
	"nxg/internal/api/rest"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/repository"
	"nxg/internal/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	svc service.EmployeeService
}

func SetupEmployeeRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.EmployeeService{
		Repo:         repository.NewEmployeeRepository(rh.DB),
		Auth:         rh.Auth,
		LoggedInUser: domain.User{}, // dummy user for now, needs to be replaced with actual logged in user
	}

	handler := EmployeeHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/employee")
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/:id", handler.GetEmployee)
	pvtRoutes.Get("/", handler.GetEmployees)
	pvtRoutes.Post("/create", handler.CreateEmployee)
	pvtRoutes.Put("/update/:id", handler.UpdateEmployee)
}

func (d *EmployeeHandler) GetEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	desgid, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	desg, err := d.svc.GetEmployeeById(desgid)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Employee fetched successfully", desg)
}

func (d *EmployeeHandler) GetEmployees(ctx *fiber.Ctx) error {
	desg, err := d.svc.GetEmployees()
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}
	return rest.SuccessResponse(ctx, "all Employees fetched successfully", desg)
}

func (d *EmployeeHandler) CreateEmployee(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	emp := dto.EmployeeCreateDto{}
	err := ctx.BodyParser(&emp)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	nemp, err := d.svc.CreateEmployee(&emp)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	user := domain.User{
		Email:     nemp.Email,
		Password:  emp.Password,
		UserType:  emp.Role,
		Status:    1,
		CreatedAt: time.Now(),
		CreatedBy: int(d.svc.LoggedInUser.ID),
	}
	_, err = d.svc.Repo.CreateUser(user)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Employees created successfully", nemp)
}

func (d *EmployeeHandler) UpdateEmployee(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	id := ctx.Params("id")
	empId, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, errors.New("invalid id format proivded in the request url"), []string{})
	}

	desg := dto.EmployeeUpdateDto{}
	err = ctx.BodyParser(&desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	udesg, err := d.svc.UpdateEmployee(int(empId), &desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Employees updated successfully", udesg)
}
