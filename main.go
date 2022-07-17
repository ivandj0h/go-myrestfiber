package main

import (
	"github.com/gofiber/fiber/v2"
	"ivandjoh.online/m/v2/routes"
)

func main() {

	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(":2022")
}
