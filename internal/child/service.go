package child

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) SeedChildren() error {

	children := []Child{
		{Name: "Kurian", Age: 10},
		{Name: "Mikayel", Age: 7},
		{Name: "Eappan", Age: 4},
		{Name: "Rahael", Age: 2},
		{Name: "Daveed", Age: 2},
	}

	for _, child := range children {

		var count int64

		s.db.Model(&Child{}).
			Where("name = ?", child.Name).
			Count(&count)

		if count == 0 {
			s.db.Create(&child)
		}
	}

	return nil
}

func (s *Service) GetAll() ([]Child, error) {

	var children []Child

	err := s.db.
		Order("total_points desc").
		Find(&children).
		Error

	return children, err
}