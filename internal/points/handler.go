package points

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

func (h *Handler) Award(c *fiber.Ctx) error {

	childID, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return fiber.ErrBadRequest
	}

	var req AwardRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.Award(
		uint(childID),
		req.Points,
		req.Reason,
	)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "points awarded",
	})
}