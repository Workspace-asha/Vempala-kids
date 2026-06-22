package reward

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	api fiber.Router,
	handler *Handler,
) {

	group := api.Group("/rewards")

	group.Get("/", handler.GetAll)
	group.Post("/:childId/:rewardId", handler.Redeem)
}
