package reward

import "time"

type Reward struct {
	ID uint `gorm:"primaryKey"`

	Title string

	Cost int

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Redemption struct {
	ID uint `gorm:"primaryKey"`

	ChildID uint

	RewardID uint

	PointsSpent int

	CreatedAt time.Time
}