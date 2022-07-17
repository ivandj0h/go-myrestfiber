package routes

import (
	"github.com/gofiber/fiber/v2"
	"ivandjoh.online/m/v2/handler"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserReadHandler)
}
