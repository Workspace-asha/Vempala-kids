package pkg

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/asha/vempala-kids/internal/child"
	"github.com/asha/vempala-kids/internal/points"
	"github.com/asha/vempala-kids/internal/reward"
	"github.com/asha/vempala-kids/internal/task"
	"github.com/asha/vempala-kids/internal/streak"
)


func Connect(dbPath string) (*gorm.DB, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, err
	}

	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&child.Child{},
		&task.Task{},
		&task.Assignment{},
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