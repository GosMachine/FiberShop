package handlers

import (
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/database/redis"
	"FiberShop/internal/transport/grpc/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

type Handle struct {
	Log    *zap.Logger
	Client *auth.Client
	Db     *postgres.Storage
	Redis  *redis.Redis
}

func New(log *zap.Logger, client *auth.Client, db *postgres.Storage, redis *redis.Redis) *Handle {
	return &Handle{Log: log, Client: client, Db: db, Redis: redis}
}

func SetCookie(name, value string, c *fiber.Ctx, expires time.Time) {
	cookie := fiber.Cookie{
		Name:    name,
		Secure:  true,
		Value:   value,
		Expires: expires,
	}
	c.Cookie(&cookie)
}

func (a *Handle) sendEmail(email string) {
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email, code, time.Minute*10)
	message := gomail.NewMessage()
	message.SetHeader("From", "support@fiber.shop")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "FiberShop")
	message.SetBody("text/plain", code)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "masterok6000@gmail.com", "atgtullwzawcfexa")
	if err := dialer.DialAndSend(message); err != nil {
		a.Log.Error("error send verification code "+email, zap.Error(err))
		return
	}
	a.Log.Info("send verification code", zap.String("email", email))
}
