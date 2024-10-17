package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterCustomer(basePath string, api fiber.Router) {
	customerGroup := api.Group(basePath)

	customerGroup.Get("/", controllers.CustomerReadList)
	customerGroup.Get("/:id", controllers.CustomerReadDetail)
	customerGroup.Post("/", controllers.CustomerCreate)
	customerGroup.Put("/:id", controllers.CustomerUpdate)
	customerGroup.Delete("/:id", controllers.CustomerDelete)
}
