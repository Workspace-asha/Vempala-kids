package points

import "time"

type History struct {
	ID uint `gorm:"primaryKey"`

	ChildID uint

	Points int

	Reason string

	CreatedAt time.Time
}

type AwardRequest struct {
	Points int    `json:"points"`
	Reason string `json:"reason"`
}