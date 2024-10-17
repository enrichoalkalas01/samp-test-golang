package controllers

import (
	"log"
	"time"

	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/enrichoalkalas01/samp-test-golang/models/migrations"
	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PengeluaranBarangRequest struct {
	TrxOutNo      string                     `json:"trx_out_no"`
	TrxOutDate    string                     `json:"trx_out_date"`
	WhsIdf        int                        `json:"whs_idf"`
	TrxOutSuppIdf int                        `json:"trx_out_supp_idf"`
	TrxOutNotes   string                     `json:"trx_out_notes"`
	Details       []PengeluaranDetailRequest `json:"details"`
}

type PengeluaranDetailRequest struct {
	ProductID int `json:"product_id"`
	QtyDus    int `json:"qty_dus"`
	QtyPcs    int `json:"qty_pcs"`
}

func PengeluaranBarangReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	db := models.DB.Model(&migrations.PengeluaranBarangHeader{})

	if searchQuery != "" {
		db = db.Where("trx_out_no LIKE ?", "%"+searchQuery+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		log.Printf("Error saat menghitung total data: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to count data",
			"status":  fiber.StatusInternalServerError,
		})
	}

	offset := (page - 1) * size
	db = db.Offset(offset).Limit(size)

	if sortBy != "" && order != "" {
		db = db.Order(sortBy + " " + order)
	} else {
		db = db.Order("trx_out_no asc")
	}

	var pengeluaranList []migrations.PengeluaranBarangHeader
	if err := db.Preload("Details").Find(&pengeluaranList).Error; err != nil {
		log.Printf("Error saat mengambil daftar Pengeluaran Barang: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch PengeluaranBarang list",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched PengeluaranBarang list",
		"status":  fiber.StatusOK,
		"data":    pengeluaranList,
		"pagination": fiber.Map{
			"search":    searchQuery,
			"page":      page,
			"size":      size,
			"total":     total,
			"order":     order,
			"sort_by":   sortBy,
			"totalPage": (total + int64(size) - 1) / int64(size),
		},
	})
}

func PengeluaranBarangReadDetail(c *fiber.Ctx) error {
	trxOutNo := c.Params("trx_out_no")
	if trxOutNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Transaction number is required",
			"status":  fiber.StatusBadRequest,
		})
	}

	var pengeluaranHeader migrations.PengeluaranBarangHeader
	if err := models.DB.Preload("Details").Where("trx_out_no = ?", trxOutNo).First(&pengeluaranHeader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Pengeluaran Barang not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching PengeluaranBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched PengeluaranBarang detail",
		"status":  fiber.StatusOK,
		"data":    pengeluaranHeader,
	})
}

func PengeluaranBarangCreate(c *fiber.Ctx) error {
	var body PengeluaranBarangRequest
	errorsMap, err := utils.ValidateStruct(c, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errorsMap,
			"status":  fiber.StatusBadRequest,
		})
	}

	layout := "2006-01-02"
	trxOutDate, err := time.Parse(layout, body.TrxOutDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid date format, expected YYYY-MM-DD",
			"status":  fiber.StatusBadRequest,
		})
	}

	tx := models.DB.Begin()

	pengeluaranHeader := migrations.PengeluaranBarangHeader{
		TrxOutNo:      body.TrxOutNo,
		TrxOutDate:    trxOutDate,
		WhsIdf:        body.WhsIdf,
		TrxOutSuppIdf: body.TrxOutSuppIdf,
		TrxOutNotes:   body.TrxOutNotes,
	}

	if err := tx.Create(&pengeluaranHeader).Error; err != nil {
		tx.Rollback()
		log.Printf("Error saat menyimpan PengeluaranBarangHeader: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create PengeluaranBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	for _, detail := range body.Details {
		pengeluaranDetail := migrations.PengeluaranBarangDetail{
			TrxOutIDF:         pengeluaranHeader.TrxOutPK,
			TrxOutDProductIdf: detail.ProductID,
			TrxOutDQtyDus:     detail.QtyDus,
			TrxOutDQtyPcs:     detail.QtyPcs,
		}

		if err := tx.Create(&pengeluaranDetail).Error; err != nil {
			tx.Rollback()
			log.Printf("Error saat menyimpan PengeluaranBarangDetail: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create PengeluaranBarangDetail",
				"status":  fiber.StatusInternalServerError,
			})
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error saat commit transaksi: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to commit transaction",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created PengeluaranBarang",
		"status":  200,
		"data":    body,
	})
}

func PengeluaranBarangUpdate(c *fiber.Ctx) error {
	trxOutNo := c.Params("trx_out_no")
	if trxOutNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Transaction number is required",
			"status":  fiber.StatusBadRequest,
		})
	}

	var body PengeluaranBarangRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	var pengeluaranHeader migrations.PengeluaranBarangHeader
	if err := models.DB.Where("trx_out_no = ?", trxOutNo).First(&pengeluaranHeader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Pengeluaran Barang not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching PengeluaranBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	layout := "2006-01-02"
	trxOutDate, err := time.Parse(layout, body.TrxOutDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid date format, expected YYYY-MM-DD",
			"status":  fiber.StatusBadRequest,
		})
	}

	pengeluaranHeader.TrxOutDate = trxOutDate
	pengeluaranHeader.WhsIdf = body.WhsIdf
	pengeluaranHeader.TrxOutSuppIdf = body.TrxOutSuppIdf
	pengeluaranHeader.TrxOutNotes = body.TrxOutNotes

	if err := models.DB.Save(&pengeluaranHeader).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update PengeluaranBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Where("trx_out_id_f = ?", pengeluaranHeader.TrxOutPK).Delete(&migrations.PengeluaranBarangDetail{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete old PengeluaranBarangDetail",
			"status":  fiber.StatusInternalServerError,
		})
	}

	for _, detail := range body.Details {
		pengeluaranDetail := migrations.PengeluaranBarangDetail{
			TrxOutIDF:         pengeluaranHeader.TrxOutPK,
			TrxOutDProductIdf: detail.ProductID,
			TrxOutDQtyDus:     detail.QtyDus,
			TrxOutDQtyPcs:     detail.QtyPcs,
		}

		if err := models.DB.Create(&pengeluaranDetail).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create PengeluaranBarangDetail",
				"status":  fiber.StatusInternalServerError,
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated PengeluaranBarang",
		"status":  fiber.StatusOK,
		"data":    body,
	})
}

func PengeluaranBarangDelete(c *fiber.Ctx) error {
	trxOutNo := c.Params("trx_out_no")
	if trxOutNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Transaction number is required",
			"status":  fiber.StatusBadRequest,
		})
	}

	var pengeluaranHeader migrations.PengeluaranBarangHeader
	if err := models.DB.Where("trx_out_no = ?", trxOutNo).First(&pengeluaranHeader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Pengeluaran Barang not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching PengeluaranBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Where("trx_out_id_f = ?", pengeluaranHeader.TrxOutPK).Delete(&migrations.PengeluaranBarangDetail{}).Error; err != nil {
		log.Printf("Error saat menghapus PengeluaranBarangDetail dengan TrxOutIDF %d: %v\n", pengeluaranHeader.TrxOutPK, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete PengeluaranBarangDetail",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Delete(&pengeluaranHeader).Error; err != nil {
		log.Printf("Error saat menghapus PengeluaranBarangHeader: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete PengeluaranBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted PengeluaranBarang",
		"status":  fiber.StatusOK,
	})
}
