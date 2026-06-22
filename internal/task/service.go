package task

import (
	"errors"
	"strings"
	"time"

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

func (s *Service) CreateTask(title string, points int) (*Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("task title is required")
	}

	if points <= 0 {
		return nil, errors.New("task points must be greater than zero")
	}

	task := Task{
		Title:    title,
		Points:   points,
		IsActive: true,
	}

	if err := s.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *Service) GetAvailableTasks(childID uint) ([]Task, error) {
	assignedTasks := s.db.
		Model(&Assignment{}).
		Select("task_id").
		Where("child_id = ?", childID)

	var tasks []Task
	err := s.db.
		Where("is_active = ?", true).
		Where("id NOT IN (?)", assignedTasks).
		Order("created_at desc").
		Find(&tasks).
		Error

	return tasks, err
}

func (s *Service) GetAssignments(childID uint) ([]AssignmentView, error) {
	var assignments []Assignment

	err := s.db.
		Preload("Task").
		Where("child_id = ?", childID).
		Order("assigned_at desc").
		Find(&assignments).
		Error
	if err != nil {
		return nil, err
	}

	views := make([]AssignmentView, 0, len(assignments))
	for _, assignment := range assignments {
		views = append(views, AssignmentView{
			ID:          assignment.ID,
			ChildID:     assignment.ChildID,
			TaskID:      assignment.TaskID,
			TaskTitle:   assignment.Task.Title,
			Points:      assignment.Task.Points,
			Status:      assignment.Status,
			AssignedAt:  assignment.AssignedAt,
			CompletedAt: assignment.CompletedAt,
		})
	}

	return views, nil
}

func (s *Service) AssignTask(childID uint, taskID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var childRecord child.Child
		if err := tx.First(&childRecord, childID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("child not found")
			}
			return err
		}

		var task Task
		if err := tx.First(&task, taskID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("task not found")
			}
			return err
		}

		now := time.Now()
		var assignment Assignment
		err := tx.Where("child_id = ? AND task_id = ?", childID, taskID).
			First(&assignment).
			Error

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			return tx.Create(&Assignment{
				ChildID:    childID,
				TaskID:     taskID,
				Status:     "assigned",
				AssignedAt: now,
			}).Error
		}

		return errors.New("task already added for this child")
	})
}

func (s *Service) CompleteTask(
	childID uint,
	taskID uint,
) error {
	completionErr := s.db.Transaction(func(tx *gorm.DB) error {
		var assignment Assignment
		if err := tx.Where("child_id = ? AND task_id = ?", childID, taskID).
			First(&assignment).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("task not assigned to child")
			}
			return err
		}

		if assignment.Status == "completed" {
			return errors.New("task already completed")
		}

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

		now := time.Now()
		assignment.Status = "completed"
		assignment.CompletedAt = &now

		if err := tx.Save(&assignment).Error; err != nil {
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

	if completionErr != nil {
		return completionErr
	}

	return nil
}