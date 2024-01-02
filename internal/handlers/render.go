package handlers

import (
	"FiberShop/internal/models"
	"FiberShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func renderTemplate(c *fiber.Ctx, tmpl string, a *Handle, data fiber.Map) error {
	var IsAuthenticated bool
	//url := c.OriginalURL()
	//ip := c.IP()
	email, token := utils.IsUserLoggedIn(c.Cookies("token"), a.Log)
	if email != "" {
		IsAuthenticated = true
	}
	if IsAuthenticated && token != "" {
		setCookie("token", token, c)
	}
	var user models.User
	if email != "" {
		var err error
		user, err = a.Db.User(email)
		if err != nil {
			a.Log.Error("error getting user", zap.Error(err))
		}
	}
	FinalData := struct {
		IsAuthenticated bool
		Balance         float64
		//Viewers        string
		Data interface{}
	}{
		IsAuthenticated: IsAuthenticated,
		Balance:         user.Balance,
		//Viewers:        viewersCount,
		Data: data,
	}
	return c.Render(tmpl, FinalData)
}
