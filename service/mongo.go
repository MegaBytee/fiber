package service

import (
	"github.com/MegaBytee/fiber/mongo"
	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
)

//default service for mongodb transaction

var DB = mongo.New(mongo.Config{
	Host:     utils.GetEnv("DB_HOST"),
	Database: utils.GetEnv("DB_NAME"),
	Username: utils.GetEnv("DB_USER"),
	Password: utils.GetEnv("DB_PASS"),
})

var Mongo = NewService("mongo").
	SetHandler(
		Handler{
			Method: utils.METHOD_GET,
			Path:   "/health",
			Func:   check_mg_health,
		})

func check_mg_health(c *fiber.Ctx) error {
	utils.SetNoCacheControl(c)
	data := map[string]any{
		"connected": DB.Connected(),
	}

	return c.JSON(data)
}
