package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/asha/vempala-kids/internal/child"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Index(c *fiber.Ctx) error {

	var children []child.Child

	if err := h.db.
		Order("total_points desc").
		Find(&children).Error; err != nil {
		return err
	}

	return c.Render("dashboard", fiber.Map{
		"Children": children,
	})
}