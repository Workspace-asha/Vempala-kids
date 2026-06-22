package task

import (
	"errors"

	"gorm.io/gorm"

	"github.com/asha/vempala-kids/internal/child"
	"github.com/asha/vempala-kids/internal/points"
	"github.com/asha/vempala-kids/internal/streak"
)

type Service struct {
	db            *gorm.DB
	streakService *streak.Service
}

func NewService(
	db *gorm.DB,
	streakService *streak.Service,
) *Service {
	return &Service{
		db:            db,
		streakService: streakService,
	}
}

func (s *Service) SeedTasks() error {

	tasks := []Task{
		{
			Title:    "Make Bed",
			Points:   10,
			IsActive: true,
		},
		{
			Title:    "Brush Teeth",
			Points:   5,
			IsActive: true,
		},
		{
			Title:    "Homework",
			Points:   20,
			IsActive: true,
		},
		{
			Title:    "Reading",
			Points:   15,
			IsActive: true,
		},
		{
			Title:    "Help Parents",
			Points:   15,
			IsActive: true,
		},
	}

	for _, task := range tasks {

		var count int64

		s.db.Model(&Task{}).
			Where("title = ?", task.Title).
			Count(&count)

		if count == 0 {
			s.db.Create(&task)
		}
	}

	return nil
}

func (s *Service) GetAll() ([]Task, error) {

	var tasks []Task

	err := s.db.
		Where("is_active = ?", true).
		Find(&tasks).
		Error

	return tasks, err
}

func (s *Service) CompleteTask(
	childID uint,
	taskID uint,
) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		var task Task

		if err := tx.First(&task, taskID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("task not found")
			}
			return err
		}

		var childRecord child.Child

		if err := tx.First(&childRecord, childID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("child not found")
			}
			return err
		}

		childRecord.TotalPoints += task.Points

		if err := tx.Save(&childRecord).Error; err != nil {
			return err
		}

		history := points.History{
			ChildID: childID,
			Points:  task.Points,
			Reason:  task.Title,
		}

		if err := tx.Create(&history).Error; err != nil {
			return err
		}

		return nil
	})
}