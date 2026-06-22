package reward

import (
	"errors"

	"gorm.io/gorm"

	"github.com/asha/vempala-kids/internal/child"
	"github.com/asha/vempala-kids/internal/points"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) SeedRewards() error {

	rewards := []Reward{
		{Title: "Ice Cream", Cost: 100, IsActive: true},
		{Title: "Chocolate", Cost: 150, IsActive: true},
		{Title: "Movie Night", Cost: 300, IsActive: true},
		{Title: "New Toy", Cost: 1000, IsActive: true},
	}

	for _, reward := range rewards {

		var count int64

		s.db.Model(&Reward{}).
			Where("title = ?", reward.Title).
			Count(&count)

		if count == 0 {
			s.db.Create(&reward)
		}
	}

	return nil
}

func (s *Service) GetAll() ([]Reward, error) {

	var rewards []Reward

	err := s.db.
		Where("is_active = ?", true).
		Find(&rewards).
		Error

	return rewards, err
}

func (s *Service) Redeem(
	childID uint,
	rewardID uint,
) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		var c child.Child
		if err := tx.First(&c, childID).Error; err != nil {
			return err
		}

		var r Reward
		if err := tx.First(&r, rewardID).Error; err != nil {
			return err
		}

		if c.TotalPoints < r.Cost {
			return errors.New("not enough points")
		}

		c.TotalPoints -= r.Cost

		if err := tx.Save(&c).Error; err != nil {
			return err
		}

		redemption := Redemption{
			ChildID:     childID,
			RewardID:    rewardID,
			PointsSpent: r.Cost,
		}

		if err := tx.Create(&redemption).Error; err != nil {
			return err
		}

		history := points.History{
			ChildID: childID,
			Points: -r.Cost,
			Reason: "Redeemed: " + r.Title,
		}

		return tx.Create(&history).Error
	})
}