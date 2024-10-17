package controllers

import (
	"log"

	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/enrichoalkalas01/samp-test-golang/models/migrations"
	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	db := models.DB.Model(&migrations.Product{})

	if searchQuery != "" {
		db = db.Where("product_name LIKE ?", "%"+searchQuery+"%")
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
		db = db.Order("product_name asc")
	}

	var productList []migrations.Product
	if err := db.Find(&productList).Error; err != nil {
		log.Printf("Error saat mengambil daftar Product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Product list",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Product list",
		"status":  fiber.StatusOK,
		"data":    productList,
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

func ProductReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var product migrations.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Product",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Product detail",
		"status":  fiber.StatusOK,
		"data":    product,
	})
}

func ProductCreate(c *fiber.Ctx) error {
	var body migrations.Product

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	if err := models.DB.Create(&body).Error; err != nil {
		log.Printf("Error saat menyimpan Product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create Product",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created Product",
		"status":  fiber.StatusOK,
		"data":    body,
	})
}

func ProductUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var body migrations.Product
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	var product migrations.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Product",
			"status":  fiber.StatusInternalServerError,
		})
	}

	product.ProductName = body.ProductName

	if err := models.DB.Save(&product).Error; err != nil {
		log.Printf("Error saat mengupdate Product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Product",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated Product",
		"status":  fiber.StatusOK,
		"data":    product,
	})
}

func ProductDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var product migrations.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Product",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Delete(&product).Error; err != nil {
		log.Printf("Error saat menghapus Product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Product",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted Product",
		"status":  fiber.StatusOK,
	})
}
