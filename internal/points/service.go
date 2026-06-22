package points

import (
	"gorm.io/gorm"

	"github.com/asha/vempala-kids/internal/child"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Award(
	childID uint,
	points int,
	reason string,
) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		var c child.Child

		if err := tx.First(&c, childID).Error; err != nil {
			return err
		}

		c.TotalPoints += points

		if err := tx.Save(&c).Error; err != nil {
			return err
		}

		history := History{
			ChildID: childID,
			Points:  points,
			Reason:  reason,
		}

		return tx.Create(&history).Error
	})
}