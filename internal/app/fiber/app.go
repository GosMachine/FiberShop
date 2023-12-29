package fiberapp

import (
	"FiberShop/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
	app *fiber.App
}

func New(log *zap.Logger) *App {
	app := fiber.New()
	//route := routes2.SetupPublicRoutes()
	middle := middleware.New(log)
	app.Use(middle.Logger, cors.New())

	return &App{app: app, log: log}
}

func (a *App) Run() {
	err := a.app.Listen(":3000")
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
