package controllers

import (
	"fmt"

	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
)

// Request Body Field & Validator ( tag )
type templateRequest struct {
	Name  string `json:"name" validate:"required,min=1"`  // Name harus diisi, min 1 karakter
	Email string `json:"email" validate:"required,email"` // Email harus valid
	Age   int    `json:"age" validate:"gte=17,lte=65"`    // Umur harus antara 17 sampai 65 tahun
}

func TemplateReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams((c)) // Parsing Automate Query Params
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfull to get Template",
		"status":  200,
		"pagination": fiber.Map{
			"search":  searchQuery,
			"page":    page,
			"size":    size,
			"order":   order,
			"sort_by": sortBy,
		},
	})
}

func TemplateReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	return c.JSON(fiber.Map{
		"message": "Successfull to get Template detail",
		"status":  200,
	})
}

func TemplateCreate(c *fiber.Ctx) error {
	var body templateRequest

	// Memvalidasi request body
	errorsMap, err := utils.ValidateStruct(c, &body)
	if err != nil {
		// Jika validasi gagal, kembalikan respons yang lebih rinci
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"errors":  errorsMap, // Mengembalikan detail kesalahan validasi
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(body)
	return c.JSON(fiber.Map{
		"message": "Successfull to create Template",
		"status":  200,
		"data":    body,
	})
}

func TemplateUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	var body templateRequest

	// Memvalidasi request body
	errorsMap, err := utils.ValidateStruct(c, &body)
	if err != nil {
		// Jika validasi gagal, kembalikan respons yang lebih rinci
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"errors":  errorsMap, // Mengembalikan detail kesalahan validasi
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(body)

	return c.JSON(fiber.Map{
		"message": "Successfull to update Template",
		"status":  200,
		"data":    body,
	})
}

func TemplateDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	return c.JSON(fiber.Map{
		"message": "Successfull to delete Template",
		"status":  200,
	})
}
