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
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"go.uber.org/zap"
	"os"
)

type App struct {
	log *zap.Logger
	app *fiber.App
}

func New(log *zap.Logger, authClient *auth.Client, db *postgres.Storage, redis *redis.Redis) *App {
	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	middle := middleware.New(log)
	goth.UseProviders(
		google.New(os.Getenv("googleClientKey"), os.Getenv("googleClientSecret"), "http://localhost:3000/auth/callback/google"),
	)
	handle := handlers.New(log, authClient, db, redis)
	app.Use(middle.Logger, cors.New())
	routes.SetupRoutes(app, handle, middle)
	app.Static("/", "./web/static")
	app.Use(handle.HandleNotFound)
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
