package utils

import "github.com/gofiber/fiber/v2"

func ServerErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Forbidden")
}
func NewFromFiber(a any, c *fiber.Ctx) int {
	err := c.BodyParser(&a)
	if err != nil {
		return -1
	}
	return 0
}

func SetNoCacheControl(c *fiber.Ctx) {
	c.Set(CACHE_CONTROL, NO_CACHE)
}
