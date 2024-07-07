package demo

import (
	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
)

func save_hello(c *fiber.Ctx) error {
	utils.SetNoCacheControl(c)
	name := c.Params("name")
	value := c.Params("value")

	data := Hello{
		Name:  name,
		Value: value,
	}
	r := HELLO.Save(data)

	return c.JSON(r)
}

func say_hello(c *fiber.Ctx) error {
	utils.SetNoCacheControl(c)
	name := c.Params("name")

	x := Hello{
		Name: name,
	}

	a := []Hello{}
	pagination := HELLO.Paginate(x.FilterID(), 1, 1, &a)

	data := map[string]any{
		"data":       a,
		"pagination": pagination,
	}
	return c.JSON(data)
}
