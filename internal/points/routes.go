package points

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	api fiber.Router,
	handler *Handler,
) {

	api.Post(
		"/children/:id/points",
		handler.Award,
	)
}