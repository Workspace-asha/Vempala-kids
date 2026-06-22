package streak

import "time"

type Streak struct {
	ID uint `gorm:"primaryKey"`

	ChildID uint

	CurrentDays int

	LastActivityDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}