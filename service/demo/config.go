package demo

import (
	"github.com/MegaBytee/fiber/service"
)

var Service = service.NewService("demo").
	SetSchema(HELLO).
	SetHandler(service.Handler{
		Method: "GET",
		Path:   "/hello/:name/:value",
		Func:   save_hello,
	}).
	SetHandler(service.Handler{
		Method: "GET",
		Path:   "/hello/:name/",
		Func:   say_hello,
	})
