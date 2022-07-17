package handler

import "github.com/gofiber/fiber/v2"

func UserReadHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "hey",
	})
}
