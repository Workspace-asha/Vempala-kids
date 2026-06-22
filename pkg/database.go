package pkg

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/asha/vempala-kids/internal/child"
	"github.com/asha/vempala-kids/internal/points"
	"github.com/asha/vempala-kids/internal/reward"
	"github.com/asha/vempala-kids/internal/task"
	"github.com/asha/vempala-kids/internal/streak"
)

func Connect() (*gorm.DB, error) {

	db, err := gorm.Open(
		sqlite.Open("data/vempala.db"),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&child.Child{},
		&task.Task{},
		&points.History{},
		&reward.Reward{},
		&reward.Redemption{},
		&streak.Streak{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}