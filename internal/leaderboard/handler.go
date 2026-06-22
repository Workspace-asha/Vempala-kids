package leaderboard

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

func (h *Handler) Get(c *fiber.Ctx) error {

	var children []child.Child

	h.db.Order("total_points desc").
		Find(&children)

	return c.JSON(children)
}