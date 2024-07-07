package server

import (
	"time"

	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/mongodb/v2"
)

func LimiterSettings(app *fiber.App, mongo_uri string) {
	//limiter
	fiber_storage := mongodb.New(mongodb.Config{
		ConnectionURI: mongo_uri,
		Database:      utils.FIBER_SERVER,
		Collection:    utils.FIBER_STORAGE,
		Reset:         false,
	})
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		Storage:    fiber_storage,
	}))
}
