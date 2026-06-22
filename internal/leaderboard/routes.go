package leaderboard

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	api fiber.Router,
	handler *Handler,
) {

	api.Get(
		"/leaderboard",
		handler.Get,
	)
}