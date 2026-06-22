package child

import "time"

type Child struct {
	ID uint `gorm:"primaryKey"`

	Name string
	Age  int

	TotalPoints int

	CreatedAt time.Time
	UpdatedAt time.Time
}