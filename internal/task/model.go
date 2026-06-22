package task

import "time"

type Task struct {
	ID uint `gorm:"primaryKey"`

	Title string

	Points int

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}