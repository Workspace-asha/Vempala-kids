package task

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	api fiber.Router,
	handler *Handler,
) {

	group := api.Group("/tasks")

	group.Get("/", handler.GetAll)

	group.Post(
	"/:childId/tasks/:taskId",
	handler.Complete,
)
}