package fiber

import (
	"fmt"
	"log"

	"github.com/MegaBytee/fiber/mongo"
	"github.com/MegaBytee/fiber/server"
	"github.com/MegaBytee/fiber/service"
	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Fiber struct {
	App      *fiber.App
	MongoDB  *mongo.Mongo
	Services []service.Service
}

func NewFiber() *Fiber {
	return &Fiber{
		App:      fiber.New(),
		MongoDB:  service.DB,
		Services: []service.Service{*service.Mongo},
	}
}
func (f *Fiber) SetService(x service.Service) *Fiber {
	f.Services = append(f.Services, x)
	return f
}

func (f *Fiber) loadServices() {

	for _, x := range f.Services {
		x.Load(f.App)
	}
}

func (f *Fiber) config() {
	f.App.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))

	//compress
	f.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	//cookie encrypt
	f.App.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))
	//caching
	server.CacheSettings(f.App, f.MongoDB.GetURI())
	//limiter
	server.LimiterSettings(f.App, f.MongoDB.GetURI())
	//metrics
	server.Metrics(f.App)

}

func (f *Fiber) Run() {
	f.config()
	if f.MongoDB.Connected() {
		//fmt.Println("MongoDB.Connected()====loadplugins")
		f.loadServices()
	}

	f.App.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("GM, from Fiber microservice... 👋!")
	})
	port := utils.GetEnv("FIBER_PORT")
	fmt.Println("port=", port)
	log.Fatal(f.App.Listen(port))

}
