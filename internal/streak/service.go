package streak

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Update(
	childID uint,
) error {

	now := time.Now()

	var streak Streak

	err := s.db.
		Where("child_id = ?", childID).
		First(&streak).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		return s.db.Create(&Streak{
			ChildID: childID,
			CurrentDays: 1,
			LastActivityDate: now,
		}).Error
	}

	last := streak.LastActivityDate

	diff := int(now.Sub(last).Hours() / 24)

	switch {

	case diff == 0:
		return nil

	case diff == 1:
		streak.CurrentDays++

	default:
		streak.CurrentDays = 1
	}

	streak.LastActivityDate = now

	return s.db.Save(&streak).Error
}