package fiberapp

import (
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/database/redis"
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"
	"FiberShop/internal/routes"
	"FiberShop/internal/transport/grpc/auth"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
	app *fiber.App
}

func New(log *zap.Logger, authClient *auth.Client, db *postgres.Storage, redis *redis.Redis) *App {
	app := fiber.New()
	middle := middleware.New(redis)
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_KEY"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:3000/auth/callback/google"),
	)
	handle := handlers.New(log, authClient, db, redis)
	app.Use(cors.New())
	app.Static("/", "./web/static")
	app.Use(logger.New(logger.ConfigDefault), healthcheck.New(), recover.New())
	routes := routes.New(app, handle, middle)
	routes.SetupRoutes()
	app.Use(handle.HandleNotFound)
	return &App{app: app, log: log}
}

func (a *App) Run(address string) {
	err := a.app.Listen(address)
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
