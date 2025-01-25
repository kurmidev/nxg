package rest

import "github.com/gofiber/fiber/v2"

func ErrorMessage(ctx *fiber.Ctx, status int, err error, data any) error {
	return ctx.Status(status).JSON(fiber.Map{
		"error":  err.Error(),
		"data":   data,
		"status": false,
	})
}

func SuccessResponse(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": message,
		"data":    data,
		"status":  true,
	})
}
