package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterTemplate(basePath string, api fiber.Router) {
	templateGroup := api.Group(basePath)

	templateGroup.Get("/", controllers.TemplateReadList)
	templateGroup.Get("/:id", controllers.TemplateReadDetail)
	templateGroup.Post("/", controllers.TemplateCreate)
	templateGroup.Put("/:id", controllers.TemplateUpdate)
	templateGroup.Delete("/:id", controllers.TemplateDelete)
}
