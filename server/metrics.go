package server

import (
	"fmt"

	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Metrics(app *fiber.App) {
	metrics_auth_key := utils.GenerateRandomTokenPlus(27, "@testing")
	fmt.Println("metrics:", metrics_auth_key)
	// Initialize default config (Assign the middleware to /metrics)
	app.Get("/metrics/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		if key == metrics_auth_key {
			return c.Next()
		} else {
			return c.SendString("404 Error")
		}

	}, monitor.New())
}
