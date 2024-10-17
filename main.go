package main

import (
	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/enrichoalkalas01/samp-test-golang/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Definisikan flag untuk migrasi dan rollback

	app := fiber.New()

	// Init DB
	models.InitDB()

	// Migrate Table If Not Exist In DB
	models.MigrateDB()

	// Rollback Table
	// models.DropDB()

	// Use All Router From Routes
	routes.RouterApp("/api/v1", app)

	app.Listen(":7000")
}
