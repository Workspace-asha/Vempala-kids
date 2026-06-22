package task

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	api fiber.Router,
	handler *Handler,
) {

	group := api.Group("/tasks")

	group.Get("/", handler.GetAll)
	group.Post("/", handler.Create)

	children := api.Group("/children")
	children.Get("/:childId/tasks/available", handler.GetAvailable)
	children.Get("/:childId/tasks", handler.GetAssignments)
	children.Post("/:childId/tasks/:taskId/assign", handler.Assign)
	children.Post("/:childId/tasks/:taskId/complete", handler.Complete)
}