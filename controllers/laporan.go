package controllers

import (
	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LaporanStokResponse struct {
	Warehouse string `json:"warehouse"`
	Product   string `json:"product"`
	QtyDus    int    `json:"qty_dus"`
	QtyPcs    int    `json:"qty_pcs"`
}

func LaporanStok(c *fiber.Ctx) error {
	var laporan []LaporanStokResponse

	// Join tables between Warehouse, Product, PenerimaanBarangDetail, and PengeluaranBarangDetail
	err := models.DB.
		Table("warehouses").
		Select("warehouses.whs_name AS warehouse, products.product_name AS product, " +
			"COALESCE(SUM(penerimaan_barang_details.trx_in_d_qty_dus), 0) - COALESCE(SUM(pengeluaran_barang_details.trx_out_d_qty_dus), 0) AS qty_dus, " +
			"COALESCE(SUM(penerimaan_barang_details.trx_in_d_qty_pcs), 0) - COALESCE(SUM(pengeluaran_barang_details.trx_out_d_qty_pcs), 0) AS qty_pcs").
		Joins("LEFT JOIN penerimaan_barang_headers ON penerimaan_barang_headers.whs_idf = warehouses.whs_pk").
		Joins("LEFT JOIN penerimaan_barang_details ON penerimaan_barang_details.trx_in_id_f = penerimaan_barang_headers.trx_in_pk").
		Joins("LEFT JOIN products ON products.product_pk = penerimaan_barang_details.trx_in_d_product_idf").
		Joins("LEFT JOIN pengeluaran_barang_headers ON pengeluaran_barang_headers.whs_idf = warehouses.whs_pk").
		Joins("LEFT JOIN pengeluaran_barang_details ON pengeluaran_barang_details.trx_out_id_f = pengeluaran_barang_headers.trx_out_pk").
		Group("warehouses.whs_name, products.product_name").
		Scan(&laporan).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No stock data found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching stock report",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched stock report",
		"status":  fiber.StatusOK,
		"data":    laporan,
	})
}
