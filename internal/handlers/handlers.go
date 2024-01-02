package handlers

import (
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/transport/grpc/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handle struct {
	Log    *zap.Logger
	Client *auth.Client
	Db     *postgres.Storage
}

func New(log *zap.Logger, client *auth.Client, db *postgres.Storage) *Handle {
	return &Handle{Log: log, Client: client, Db: db}
}

func setCookie(name, value string, c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:   name,
		Secure: true,
		Value:  value,
	}
	c.Cookie(&cookie)
}
