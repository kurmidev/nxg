package handlers

import (
	"net/http"
	"nxg/internal/api/rest"
	"nxg/internal/domain"
	"nxg/internal/dto"
	"nxg/internal/repository"
	"nxg/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DesignationHandler struct {
	svc service.DesignationService
}

func SetupDesignationRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.DesignationService{
		Repo:         repository.NewDesignationRepository(rh.DB),
		Auth:         rh.Auth,
		LoggedInUser: domain.User{}, // dummy user for now, needs to be replaced with actual logged in user
	}

	handler := DesignationHandler{
		svc: svc,
	}

	pvtRoutes := app.Group("designation", rh.Auth.Authorize)
	pvtRoutes.Get("/:id", handler.GetDesignation)
	pvtRoutes.Get("/", handler.GetDesignations)
	pvtRoutes.Post("/create", handler.CreateDesignation)
	pvtRoutes.Put("/update/:id", handler.UpdateDesignation)
}

func (d *DesignationHandler) GetDesignation(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	desgid, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	desg, err := d.svc.GetDesignationById(desgid)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "designation fetched successfully", desg)
}

func (d *DesignationHandler) GetDesignations(ctx *fiber.Ctx) error {
	desg, err := d.svc.GetDesignations()
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}
	return rest.SuccessResponse(ctx, "all designations fetched successfully", desg)
}

func (d *DesignationHandler) CreateDesignation(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	desg := dto.DesignationCreateDto{}
	err := ctx.BodyParser(&desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	ndesg, err := d.svc.CreateDesignation(&desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	return rest.SuccessResponse(ctx, "designations created successfully", ndesg)
}

func (d *DesignationHandler) UpdateDesignation(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	id := ctx.Params("id")

	desgId, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	desg := dto.DesignationUpdateDto{}
	err = ctx.BodyParser(&desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	udesg, err := d.svc.UpdateDesignation(int(desgId), &desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "designations updated successfully", udesg)
}
