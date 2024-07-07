package server

import (
	"time"

	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/storage/mongodb/v2"
)

func CacheSettings(app *fiber.App, mongo_uri string) {
	//limiter
	fiber_cache := mongodb.New(mongodb.Config{
		ConnectionURI: mongo_uri,
		Database:      utils.FIBER_SERVER,
		Collection:    utils.FIBER_CACHE,
		Reset:         true,
	})

	//cache
	app.Use(cache.New(cache.Config{
		Storage: fiber_cache,
		Next: func(c *fiber.Ctx) bool {
			return c.GetRespHeader(utils.CACHE_CONTROL) == utils.NO_CACHE
		},
		Expiration:   10 * time.Minute,
		CacheControl: true,
	}))
}
