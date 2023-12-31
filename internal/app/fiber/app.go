package fiberapp

import (
	"FiberShop/internal/middleware"
	"FiberShop/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
	app *fiber.App
}

func New(log *zap.Logger) *App {
	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Static("/", "./web/static")
	middle := middleware.New(log)
	app.Use(middle.Logger, cors.New())
	routes.SetupPublicRoutes(app)
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
