package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterLaporan(basePath string, api fiber.Router) {
	LaporanGroup := api.Group(basePath)

	LaporanGroup.Get("/", controllers.LaporanStok)
}
