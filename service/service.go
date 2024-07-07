package service

import (
	"fmt"
	"strings"

	"github.com/MegaBytee/fiber/mongo"
	"github.com/MegaBytee/fiber/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Name   string
	Path   string
	Method string
	Func   func(c *fiber.Ctx) error
}

type Service struct {
	Name     string
	Mongo    *mongo.Mongo
	Schemas  []mongo.Schema
	Handlers []Handler
}

func NewService(name string) *Service {
	return &Service{
		Name:  name,
		Mongo: DB,
	}
}

func (s *Service) SetHandler(x Handler) *Service {
	s.Handlers = append(s.Handlers, x)
	return s
}

func (s *Service) SetSchema(schema *mongo.Schema) *Service {
	schema.SetMongo(s.Mongo)
	s.Schemas = append(s.Schemas, *schema)
	//fmt.Println("schema>mongo:", schema.Mongo)
	return s
}

func (s *Service) SetupSchemas() {
	for _, x := range s.Schemas {
		x.CreateIndex()
	}
}

func (s *Service) Load(app fiber.Router) int {
	s.SetupSchemas()
	api := app.Group("/" + strings.ToLower(s.Name))

	for _, x := range s.Handlers {
		switch x.Method {
		case utils.METHOD_GET:
			api.Get(x.Path, x.Func)

		case utils.METHOD_POST:
			api.Post(x.Path, x.Func)

		default:
			fmt.Println("error in setupHandlers = -1")
			return -1

		}
	}

	return 0

}
