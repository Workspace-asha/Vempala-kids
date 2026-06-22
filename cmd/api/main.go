package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/asha/vempala-kids/internal/child"
	"github.com/asha/vempala-kids/internal/leaderboard"
	"github.com/asha/vempala-kids/internal/points"
	"github.com/asha/vempala-kids/internal/reward"
	"github.com/asha/vempala-kids/internal/streak"
	"github.com/asha/vempala-kids/internal/task"
	"github.com/asha/vempala-kids/pkg"
)

func main() {

	// Database
	db, err := pkg.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Services
	childService := child.NewService(db)

	streakService := streak.NewService(db)

	taskService := task.NewService(
		db,
		streakService,
	)

	pointsService := points.NewService(db)

	rewardService := reward.NewService(db)

	// Seed Data
	if err := childService.SeedChildren(); err != nil {
		log.Fatal(err)
	}

	if err := taskService.SeedTasks(); err != nil {
		log.Fatal(err)
	}

	if err := rewardService.SeedRewards(); err != nil {
		log.Fatal(err)
	}

	// Handlers
	childHandler := child.NewHandler(childService)

	taskHandler := task.NewHandler(taskService)

	pointsHandler := points.NewHandler(pointsService)

	rewardHandler := reward.NewHandler(rewardService)

	leaderboardHandler := leaderboard.NewHandler(db)

	// Fiber App
	app := fiber.New()

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app":    "Vempala Kids",
			"status": "UP",
		})
	})

	// API v1
	api := app.Group("/api/v1")

	// Routes
	child.RegisterRoutes(
		api,
		childHandler,
	)

	task.RegisterRoutes(
		api,
		taskHandler,
	)

	points.RegisterRoutes(
		api,
		pointsHandler,
	)

	reward.RegisterRoutes(
		api,
		rewardHandler,
	)

	leaderboard.RegisterRoutes(
		api,
		leaderboardHandler,
	)

	log.Println("🚀 Vempala Kids started on :8080")

	log.Fatal(
		app.Listen(":8080"),
	)
}