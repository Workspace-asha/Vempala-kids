package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"

	"github.com/asha/vempala-kids/internal/child"
	"github.com/asha/vempala-kids/internal/dashboard"
	"github.com/asha/vempala-kids/internal/leaderboard"
	"github.com/asha/vempala-kids/internal/points"
	"github.com/asha/vempala-kids/internal/reward"
	"github.com/asha/vempala-kids/internal/streak"
	"github.com/asha/vempala-kids/internal/task"
	"github.com/asha/vempala-kids/pkg"
)

func main() {
	cfg := pkg.LoadConfig()

	// Database
	db, err := pkg.Connect(cfg.DBPath)
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

	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./web/static")

	renderPage := func(name string) fiber.Handler {
		return func(c *fiber.Ctx) error {
			return c.Render(name, fiber.Map{})
		}
	}

	app.Get("/leaderboard", renderPage("leaderboard"))
	app.Get("/rewards", renderPage("rewards"))

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

	dashboardHandler := dashboard.NewHandler(db)

	dashboard.RegisterRoutes(
		app,
		dashboardHandler,
	)
	log.Printf("🚀 Vempala Kids started on :%s", cfg.ServerPort)

	log.Fatal(
		app.Listen(":" + cfg.ServerPort),
	)
}