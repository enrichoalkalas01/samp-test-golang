package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouterApp(basePath string, app *fiber.App) {
	api := app.Group(basePath)

	RouterPenerimaanBarang("/penerimaan-barang", api)
	RouterPengeluaranBarang("/pengeluaran-barang", api)
	RouterCustomer("/customer", api)
	RouterProduct("/product", api)
	RouterWarehouse("/warehouse", api)
	RouterSupplier("/supplier", api)
	RouterTemplate("/template", api)
	RouterLaporan("/laporan-stock", api)
}
