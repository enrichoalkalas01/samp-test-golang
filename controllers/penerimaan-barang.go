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

type PenerimaanBarangRequest struct {
	TrxInNo      string          `json:"trx_in_no"`
	TrxInDate    string          `json:"trx_in_date"`
	WhsIdf       int             `json:"whs_idf"`
	TrxInSuppIdf int             `json:"trx_in_supp_idf"`
	TrxInNotes   string          `json:"trx_in_notes"`
	Details      []DetailRequest `json:"details"`
}

type DetailRequest struct {
	ProductID int `json:"product_id"`
	QtyDus    int `json:"qty_dus"`
	QtyPcs    int `json:"qty_pcs"`
}

func PenerimaanBarangReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	db := models.DB.Model(&migrations.PenerimaanBarangHeader{})

	if searchQuery != "" {
		db = db.Where("trx_in_no LIKE ?", "%"+searchQuery+"%")
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
		db = db.Order("trx_in_no asc")
	}

	var penerimaanList []migrations.PenerimaanBarangHeader
	if err := db.Preload("Details").Find(&penerimaanList).Error; err != nil {
		log.Printf("Error saat mengambil daftar Penerimaan Barang: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch PenerimaanBarang list",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched PenerimaanBarang list",
		"status":  fiber.StatusOK,
		"data":    penerimaanList,
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

func PenerimaanBarangReadDetail(c *fiber.Ctx) error {
	trxInNo := c.Params("trx_in_no")
	if trxInNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Transaction number is required",
			"status":  fiber.StatusBadRequest,
		})
	}

	var penerimaanHeader migrations.PenerimaanBarangHeader
	if err := models.DB.Preload("Details").Where("trx_in_no = ?", trxInNo).First(&penerimaanHeader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Penerimaan Barang not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching PenerimaanBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched PenerimaanBarang detail",
		"status":  200,
		"data":    penerimaanHeader,
	})
}

func PenerimaanBarangCreate(c *fiber.Ctx) error {
	var body PenerimaanBarangRequest
	errorsMap, err := utils.ValidateStruct(c, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errorsMap,
			"status":  fiber.StatusBadRequest,
		})
	}

	layout := "2006-01-02"
	trxInDate, err := time.Parse(layout, body.TrxInDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid date format, expected YYYY-MM-DD",
			"status":  fiber.StatusBadRequest,
		})
	}

	tx := models.DB.Begin()

	penerimaanHeader := migrations.PenerimaanBarangHeader{
		TrxInNo:      body.TrxInNo,
		TrxInDate:    trxInDate,
		WhsIdf:       body.WhsIdf,
		TrxInSuppIdf: body.TrxInSuppIdf,
		TrxInNotes:   body.TrxInNotes,
	}

	if err := tx.Create(&penerimaanHeader).Error; err != nil {
		tx.Rollback()
		log.Printf("Error saat menyimpan PenerimaanBarangHeader: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create PenerimaanBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	for _, detail := range body.Details {
		penerimaanDetail := migrations.PenerimaanBarangDetail{
			TrxInIDF:         penerimaanHeader.TrxInPK,
			TrxInDProductIdf: detail.ProductID,
			TrxInDQtyDus:     detail.QtyDus,
			TrxInDQtyPcs:     detail.QtyPcs,
		}

		if err := tx.Create(&penerimaanDetail).Error; err != nil {
			tx.Rollback()
			log.Printf("Error saat menyimpan PenerimaanBarangDetail: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create PenerimaanBarangDetail",
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
		"message": "Successfully created PenerimaanBarang",
		"status":  200,
		"data":    body,
	})
}

func PenerimaanBarangUpdate(c *fiber.Ctx) error {
	trxInNo := c.Params("trx_in_no")
	if trxInNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Transaction number is required",
			"status":  fiber.StatusBadRequest,
		})
	}

	var body PenerimaanBarangRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	var penerimaanHeader migrations.PenerimaanBarangHeader
	if err := models.DB.Where("trx_in_no = ?", trxInNo).First(&penerimaanHeader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Penerimaan Barang not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching PenerimaanBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	layout := "2006-01-02"
	trxInDate, err := time.Parse(layout, body.TrxInDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid date format, expected YYYY-MM-DD",
			"status":  fiber.StatusBadRequest,
		})
	}

	penerimaanHeader.TrxInDate = trxInDate
	penerimaanHeader.WhsIdf = body.WhsIdf
	penerimaanHeader.TrxInSuppIdf = body.TrxInSuppIdf
	penerimaanHeader.TrxInNotes = body.TrxInNotes

	if err := models.DB.Save(&penerimaanHeader).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update PenerimaanBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Where("trx_in_id_f = ?", penerimaanHeader.TrxInPK).Delete(&migrations.PenerimaanBarangDetail{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete old PenerimaanBarangDetail",
			"status":  fiber.StatusInternalServerError,
		})
	}

	for _, detail := range body.Details {
		penerimaanDetail := migrations.PenerimaanBarangDetail{
			TrxInIDF:         penerimaanHeader.TrxInPK,
			TrxInDProductIdf: detail.ProductID,
			TrxInDQtyDus:     detail.QtyDus,
			TrxInDQtyPcs:     detail.QtyPcs,
		}

		if err := models.DB.Create(&penerimaanDetail).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create PenerimaanBarangDetail",
				"status":  fiber.StatusInternalServerError,
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated PenerimaanBarang",
		"status":  fiber.StatusOK,
		"data":    body,
	})
}

func PenerimaanBarangDelete(c *fiber.Ctx) error {
	trxInNo := c.Params("trx_in_no")
	if trxInNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Transaction number is required",
			"status":  fiber.StatusBadRequest,
		})
	}

	var penerimaanHeader migrations.PenerimaanBarangHeader
	if err := models.DB.Where("trx_in_no = ?", trxInNo).First(&penerimaanHeader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Penerimaan Barang not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching PenerimaanBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Where("trx_in_id_f = ?", penerimaanHeader.TrxInPK).Delete(&migrations.PenerimaanBarangDetail{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete PenerimaanBarangDetail",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Delete(&penerimaanHeader).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete PenerimaanBarangHeader",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted PenerimaanBarang",
		"status":  fiber.StatusOK,
	})
}
