package controllers

import (
	"log"

	"github.com/enrichoalkalas01/samp-test-golang/models"
	"github.com/enrichoalkalas01/samp-test-golang/models/migrations"
	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CustomerReadList(c *fiber.Ctx) error {

	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	db := models.DB.Model(&migrations.Customer{})

	if searchQuery != "" {
		db = db.Where("customer_name LIKE ?", "%"+searchQuery+"%")
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

		db = db.Order("customer_name asc")
	}

	var customerList []migrations.Customer
	if err := db.Find(&customerList).Error; err != nil {
		log.Printf("Error saat mengambil daftar Customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Customer list",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Customer list",
		"status":  fiber.StatusOK,
		"data":    customerList,
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

func CustomerReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var customer migrations.Customer
	if err := models.DB.First(&customer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Customer not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Customer",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched Customer detail",
		"status":  fiber.StatusOK,
		"data":    customer,
	})
}

func CustomerCreate(c *fiber.Ctx) error {
	var body migrations.Customer

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	if err := models.DB.Create(&body).Error; err != nil {
		log.Printf("Error saat menyimpan Customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create Customer",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created Customer",
		"status":  fiber.StatusOK,
		"data":    body,
	})
}

func CustomerUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var body migrations.Customer
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  fiber.StatusBadRequest,
		})
	}

	var customer migrations.Customer
	if err := models.DB.First(&customer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Customer not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Customer",
			"status":  fiber.StatusInternalServerError,
		})
	}

	customer.CustomerName = body.CustomerName

	if err := models.DB.Save(&customer).Error; err != nil {
		log.Printf("Error saat mengupdate Customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update Customer",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated Customer",
		"status":  fiber.StatusOK,
		"data":    customer,
	})
}

func CustomerDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	var customer migrations.Customer
	if err := models.DB.First(&customer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Customer not found",
				"status":  fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching Customer",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := models.DB.Delete(&customer).Error; err != nil {
		log.Printf("Error saat menghapus Customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete Customer",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted Customer",
		"status":  fiber.StatusOK,
	})
}
