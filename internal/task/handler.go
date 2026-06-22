package task

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {

	tasks, err := h.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(tasks)
}

func (h *Handler) Complete(c *fiber.Ctx) error {

	childID, _ := strconv.Atoi(
		c.Params("childId"),
	)

	taskID, _ := strconv.Atoi(
		c.Params("taskId"),
	)

	err := h.service.CompleteTask(
		uint(childID),
		uint(taskID),
	)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "task completed",
	})
}