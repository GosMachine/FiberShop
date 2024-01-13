package handlers

import (
	"FiberShop/internal/models"
	"FiberShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) renderTemplate(c *fiber.Ctx, tmpl string, data fiber.Map) error {
	var isAuthenticated bool
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	var user models.User
	if email != "" {
		isAuthenticated = true
		var err error
		user, err = a.Redis.GetUserCache(email)
		if err != nil {
			a.Log.Error("error getting user", zap.Error(err))
		}
	}

	url := c.OriginalURL()
	ip := c.IP()
	a.Redis.IncrementViewCounter(url, ip+":"+url)
	viewers := a.Redis.Client.Get(a.Redis.Ctx, "viewers:"+url).Val()

	FinalData := struct {
		IsAuthenticated bool
		EmailVerified   bool
		Balance         float64
		Email           string
		Viewers         string
		Data            interface{}
	}{
		IsAuthenticated: isAuthenticated,
		EmailVerified:   user.EmailVerified,
		Balance:         user.Balance,
		Email:           email,
		Viewers:         viewers,
		Data:            data,
	}
	return c.Render(tmpl, FinalData)
}
