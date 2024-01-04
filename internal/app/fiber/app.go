package fiberapp

import (
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/database/redis"
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"
	"FiberShop/internal/routes"
	"FiberShop/internal/transport/grpc/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
	app *fiber.App
}

func New(log *zap.Logger, authClient *auth.Client, db *postgres.Storage, redis *redis.Redis) *App {
	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	middle := middleware.New(log)
	app.Use(middle.Logger, cors.New())
	app.Static("/", "./web/static")
	handle := handlers.New(log, authClient, db, redis)
	routes.SetupRoutes(app, handle, middle)
	return &App{app: app, log: log}
}

func (a *App) Run(adress string) {
	err := a.app.Listen(adress)
	if err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	err := a.app.Shutdown()
	if err != nil {
		a.log.Error("", zap.Error(err))
	}
}
