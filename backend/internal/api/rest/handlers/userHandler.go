package handlers

import (
	"net/http"
	"nxg/internal/api/rest"
	"nxg/internal/dto"
	"nxg/internal/repository"
	"nxg/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUsersRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}
	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/users")

	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	//private end points
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/verify", handler.GetVerification)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Get("/profile", handler.Profile)
	pvtRoutes.Post("/profile", handler.CreateProfile)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
			"error":   err,
		})
	}
	token, err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error registering user",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Register user",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
			"error":   err,
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error logging in user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login user",
		"token":   token,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Verify",
	})
}

func (h *UserHandler) Profile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Profile",
		"data":    user,
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "CreateProfiler",
	})
}

func (h *UserHandler) GetVerification(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "GetVerification",
	})
}
