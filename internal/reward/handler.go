package reward

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
	rewards, err := h.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(rewards)
}

func (h *Handler) Redeem(c *fiber.Ctx) error {
	childID, err := strconv.Atoi(c.Params("childId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	rewardID, err := strconv.Atoi(c.Params("rewardId"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.Redeem(uint(childID), uint(rewardID)); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "reward redeemed"})
}
