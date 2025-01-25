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

	"github.com/gofiber/fiber/v2"
)

type CompanyHandler struct {
	svc service.CompanyService
}

func SetupCompanyRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.CompanyService{
		Repo:         repository.NewCompanyRepository(rh.DB),
		Auth:         rh.Auth,
		LoggedInUser: domain.User{}, // dummy user for now, needs to be replaced with actual logged in user
	}

	handler := CompanyHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/company")
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/:id", handler.GetCompany)
	pvtRoutes.Get("/", handler.GetCompanys)
	pvtRoutes.Post("/create", handler.CreateCompany)
	pvtRoutes.Put("/update/:id", handler.UpdateCompany)
}

func (d *CompanyHandler) GetCompany(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	desgid, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	desg, err := d.svc.GetCompanyById(desgid)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Company fetched successfully", desg)
}

func (d *CompanyHandler) GetCompanys(ctx *fiber.Ctx) error {
	desg, err := d.svc.Repo.GetAllCompany()
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}
	return rest.SuccessResponse(ctx, "all Companys fetched successfully", desg)
}

func (d *CompanyHandler) CreateCompany(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	comp := dto.CompanyCreateDto{}
	err := ctx.BodyParser(&comp)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	ndesg, err := d.svc.CreateCompany(&comp)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Company created successfully", ndesg)
}

func (d *CompanyHandler) UpdateCompany(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	id := ctx.Params("id")
	compId, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, errors.New("Invalid id format proivded in the request url"), []string{})
	}
	comp := dto.CompanyUpdateDto{}
	err = ctx.BodyParser(&comp)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	ucomp, err := d.svc.UpdateCompany(int(compId), &comp)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Company details updated successfully", ucomp)
}
