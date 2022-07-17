package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"ivandjoh.online/m/v2/db"
	"ivandjoh.online/m/v2/model/entity"
	"ivandjoh.online/m/v2/model/request"
)

func UserGetAllHandler(ctx *fiber.Ctx) error {

	var users []entity.User

	// This code is used to see the SQL query in the console
	err := db.DB.Debug().Find(&users).Error

	// This code isn't appear SQL query in the console
	// err := db.DB.Find(&users).Error

	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(users)
}

func UserCreateHandler(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}

	errorCreateUser := db.DB.Create(&newUser).Error

	if errorCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User created",
		"data":    newUser,
	})
}
