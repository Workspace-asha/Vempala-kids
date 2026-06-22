package child

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {

	children, err := h.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(children)
}