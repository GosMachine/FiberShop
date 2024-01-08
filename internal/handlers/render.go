package handlers

import (
	"FiberShop/internal/models"
	"FiberShop/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

func (a *Handle) renderTemplate(c *fiber.Ctx, tmpl string, data fiber.Map) error {
	//url := c.OriginalURL()
	//ip := c.IP()
	timeStart := time.Now()
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
	FinalData := struct {
		IsAuthenticated bool
		Balance         float64
		//Viewers        string
		Data interface{}
	}{
		IsAuthenticated: isAuthenticated,
		Balance:         user.Balance,
		//Viewers:        viewersCount,
		Data: data,
	}
	fmt.Println(time.Since(timeStart))
	return c.Render(tmpl, FinalData)
}
