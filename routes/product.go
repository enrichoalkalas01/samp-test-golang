package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterProduct(basePath string, api fiber.Router) {
	productGroup := api.Group(basePath)

	productGroup.Get("/", controllers.ProductReadList)
	productGroup.Get("/:id", controllers.ProductReadDetail)
	productGroup.Post("/", controllers.ProductCreate)
	productGroup.Put("/:id", controllers.ProductUpdate)
	productGroup.Delete("/:id", controllers.ProductDelete)
}
