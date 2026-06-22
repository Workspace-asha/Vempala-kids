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

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Points int    `json:"points"`
}

type Assignment struct {
	ID uint `gorm:"primaryKey"`

	ChildID uint
	TaskID  uint

	Status string

	AssignedAt  time.Time
	CompletedAt *time.Time

	Task Task `gorm:"foreignKey:TaskID"`
}

type AssignmentView struct {
	ID uint `json:"id"`

	ChildID uint `json:"child_id"`
	TaskID  uint `json:"task_id"`

	TaskTitle string `json:"task_title"`
	Points    int    `json:"points"`
	Status    string `json:"status"`

	AssignedAt  time.Time  `json:"assigned_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}