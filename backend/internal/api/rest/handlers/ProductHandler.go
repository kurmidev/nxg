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

type ProductHandler struct {
	svc service.ProductService
}

func SetupProductRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.ProductService{
		Repo:         repository.NewProductRepository(rh.DB),
		Auth:         rh.Auth,
		LoggedInUser: domain.User{}, // dummy user for now, needs to be replaced with actual logged in user
	}

	handler := ProductHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/Product")
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/:id", handler.GetProduct)
	pvtRoutes.Get("/", handler.GetProducts)
	pvtRoutes.Post("/create", handler.CreateProduct)
	pvtRoutes.Put("/update/:id", handler.UpdateProduct)
}

func (d *ProductHandler) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	desgid, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	desg, err := d.svc.GetProductById(desgid)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Product fetched successfully", desg)
}

func (d *ProductHandler) GetProducts(ctx *fiber.Ctx) error {
	desg, err := d.svc.GetProducts()
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}
	return rest.SuccessResponse(ctx, "all Products fetched successfully", desg)
}

func (d *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	desg := dto.ProductCreateDto{}
	err := ctx.BodyParser(&desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	ndesg, err := d.svc.CreateProduct(&desg)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Products created successfully", ndesg)
}

func (d *ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	d.svc.LoggedInUser = d.svc.Auth.GetCurrentUser(ctx)
	id := ctx.Params("id")
	prodId, err := strconv.Atoi(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, errors.New("Invalid id format proivded in the request url"), []string{})
	}

	prod := dto.ProductUpdateDto{}
	err = ctx.BodyParser(&prod)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusBadRequest, err, []string{})
	}
	udesg, err := d.svc.UpdateProduct(int(prodId), &prod)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err, []string{})
	}

	return rest.SuccessResponse(ctx, "Products updated successfully", udesg)
}
