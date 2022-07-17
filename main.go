package main

import (
	"github.com/gofiber/fiber/v2"
	"ivandjoh.online/m/v2/db"
	"ivandjoh.online/m/v2/db/migration"
	"ivandjoh.online/m/v2/routes"
)

func main() {

	// Initialize DB
	db.DatabaseInit()

	// Run migration
	migration.RunMigration()

	// FiberApp Initialize
	app := fiber.New()

	// Routes
	routes.RouteInit(app)

	// Listen on port 2022
	app.Listen(":2022")
}
