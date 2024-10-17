package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterSupplier(basePath string, api fiber.Router) {
	supplierGroup := api.Group(basePath)

	supplierGroup.Get("/", controllers.SupplierReadList)
	supplierGroup.Get("/:id", controllers.SupplierReadDetail)
	supplierGroup.Post("/", controllers.SupplierCreate)
	supplierGroup.Put("/:id", controllers.SupplierUpdate)
	supplierGroup.Delete("/:id", controllers.SupplierDelete)
}
