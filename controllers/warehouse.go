package controllers

import (
	"log"

	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/enrichoalkalas01/samp-test-golang/models/migrations"
	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func WarehouseReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	db := models.DB.Model(&migrations.Warehouse{})

	if searchQuery != "" {
		db = db.Where("whs_name LIKE ?", "%"+searchQuery+"%")
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
		db = db.Order("whs_name asc")
	}

	var warehouseList []migrations.Warehouse
	if err := db.Find(&warehouseList).Error; err != nil {
		log.Printf("Error saat mengambil daftar Warehouse: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Warehouse list",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Warehouse list",
		"status":  fiber.StatusOK,
		"data":    warehouseList,
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

func WarehouseReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var warehouse migrations.Warehouse
	if err := models.DB.First(&warehouse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Warehouse not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Warehouse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Warehouse detail",
		"status":  fiber.StatusOK,
		"data":    warehouse,
	})
}

func WarehouseCreate(c *fiber.Ctx) error {
	var body migrations.Warehouse

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	if err := models.DB.Create(&body).Error; err != nil {
		log.Printf("Error saat menyimpan Warehouse: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create Warehouse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created Warehouse",
		"status":  fiber.StatusOK,
		"data":    body,
	})
}

func WarehouseUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var body migrations.Warehouse
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	var warehouse migrations.Warehouse
	if err := models.DB.First(&warehouse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Warehouse not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Warehouse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	warehouse.WhsName = body.WhsName

	if err := models.DB.Save(&warehouse).Error; err != nil {
		log.Printf("Error saat mengupdate Warehouse: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Warehouse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated Warehouse",
		"status":  fiber.StatusOK,
		"data":    warehouse,
	})
}

func WarehouseDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var warehouse migrations.Warehouse
	if err := models.DB.First(&warehouse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Warehouse not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Warehouse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Delete(&warehouse).Error; err != nil {
		log.Printf("Error saat menghapus Warehouse: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Warehouse",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted Warehouse",
		"status":  fiber.StatusOK,
	})
}
