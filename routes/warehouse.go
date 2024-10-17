package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterWarehouse(basePath string, api fiber.Router) {
	warehouseGroup := api.Group(basePath)

	warehouseGroup.Get("/", controllers.WarehouseReadList)
	warehouseGroup.Get("/:id", controllers.WarehouseReadDetail)
	warehouseGroup.Post("/", controllers.WarehouseCreate)
	warehouseGroup.Put("/:id", controllers.WarehouseUpdate)
	warehouseGroup.Delete("/:id", controllers.WarehouseDelete)
}
