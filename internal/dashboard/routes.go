package dashboard

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	app *fiber.App,
	handler *Handler,
) {
	app.Get("/", handler.Index)
}