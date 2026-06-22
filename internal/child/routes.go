package child

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	api fiber.Router,
	handler *Handler,
) {

	children := api.Group("/children")

	children.Get("/", handler.GetAll)
}