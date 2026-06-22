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

func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	task, err := h.service.CreateTask(req.Title, req.Points)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *Handler) GetAvailable(c *fiber.Ctx) error {
	childID, err := strconv.Atoi(c.Params("childId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	tasks, err := h.service.GetAvailableTasks(uint(childID))
	if err != nil {
		return err
	}

	return c.JSON(tasks)
}

func (h *Handler) GetAssignments(c *fiber.Ctx) error {
	childID, err := strconv.Atoi(c.Params("childId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	assignments, err := h.service.GetAssignments(uint(childID))
	if err != nil {
		return err
	}

	return c.JSON(assignments)
}

func (h *Handler) Assign(c *fiber.Ctx) error {
	childID, err := strconv.Atoi(c.Params("childId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	taskID, err := strconv.Atoi(c.Params("taskId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.AssignTask(uint(childID), uint(taskID)); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "task assigned",
	})
}

func (h *Handler) Complete(c *fiber.Ctx) error {

	childID, err := strconv.Atoi(c.Params("childId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	taskID, err := strconv.Atoi(c.Params("taskId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.CompleteTask(
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