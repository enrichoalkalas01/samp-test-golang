package controllers

import (
	"log"

	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/enrichoalkalas01/samp-test-golang/models/migrations"
	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SupplierReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	db := models.DB.Model(&migrations.Supplier{})

	if searchQuery != "" {
		db = db.Where("supplier_name LIKE ?", "%"+searchQuery+"%")
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
		db = db.Order("supplier_name asc")
	}

	var supplierList []migrations.Supplier
	if err := db.Find(&supplierList).Error; err != nil {
		log.Printf("Error saat mengambil daftar Supplier: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Supplier list",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Supplier list",
		"status":  fiber.StatusOK,
		"data":    supplierList,
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

func SupplierReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var supplier migrations.Supplier
	if err := models.DB.First(&supplier, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Supplier not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Supplier",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Supplier detail",
		"status":  fiber.StatusOK,
		"data":    supplier,
	})
}

func SupplierCreate(c *fiber.Ctx) error {
	var body migrations.Supplier

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	if err := models.DB.Create(&body).Error; err != nil {
		log.Printf("Error saat menyimpan Supplier: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create Supplier",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created Supplier",
		"status":  fiber.StatusOK,
		"data":    body,
	})
}

func SupplierUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var body migrations.Supplier
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	var supplier migrations.Supplier
	if err := models.DB.First(&supplier, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Supplier not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Supplier",
			"status":  fiber.StatusInternalServerError,
		})
	}

	supplier.SupplierName = body.SupplierName

	if err := models.DB.Save(&supplier).Error; err != nil {
		log.Printf("Error saat mengupdate Supplier: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Supplier",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated Supplier",
		"status":  fiber.StatusOK,
		"data":    supplier,
	})
}

func SupplierDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var supplier migrations.Supplier
	if err := models.DB.First(&supplier, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Supplier not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Supplier",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Delete(&supplier).Error; err != nil {
		log.Printf("Error saat menghapus Supplier: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Supplier",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted Supplier",
		"status":  fiber.StatusOK,
	})
}
