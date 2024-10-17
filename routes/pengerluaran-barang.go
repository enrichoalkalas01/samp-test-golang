package routes

import (
	"github.com/enrichoalkalas01/samp-test-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterPengeluaranBarang(basePath string, api fiber.Router) {
	pengeluaranBarangGroup := api.Group(basePath)

	pengeluaranBarangGroup.Get("/", controllers.PengeluaranBarangReadList)
	pengeluaranBarangGroup.Get("/:trx_out_no", controllers.PengeluaranBarangReadDetail)
	pengeluaranBarangGroup.Post("/", controllers.PengeluaranBarangCreate)
	pengeluaranBarangGroup.Put("/:trx_out_no", controllers.PengeluaranBarangUpdate)
	pengeluaranBarangGroup.Delete("/:trx_out_no", controllers.PengeluaranBarangDelete)
}
