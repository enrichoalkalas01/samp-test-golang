package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterPenerimaanBarang(basePath string, api fiber.Router) {
	penerimaanBarangGroup := api.Group(basePath)

	penerimaanBarangGroup.Get("/", controllers.PenerimaanBarangReadList)
	penerimaanBarangGroup.Get("/:trx_in_no", controllers.PenerimaanBarangReadDetail)
	penerimaanBarangGroup.Post("/", controllers.PenerimaanBarangCreate)
	penerimaanBarangGroup.Put("/:trx_in_no", controllers.PenerimaanBarangUpdate)
	penerimaanBarangGroup.Delete("/:trx_in_no", controllers.PenerimaanBarangDelete)
}
