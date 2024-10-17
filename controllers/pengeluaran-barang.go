package controllers

import (
	"fmt"

	"github.com/enrichoalkalas01/samp-test-golang/utils"
	"github.com/gofiber/fiber/v2"
)

// Request Body Field & Validator ( tag )
type PengeluranBarangRequest struct {
	NomorTransaksi    string  `json:"NomorTransaksi" validate:"required,min=1"`
	TanggalPengeluran string  `json:"TanggalPengeluran" validate:"required,min=1"`
	Gudang            string  `json:"Gudang" validate:"required,min=1"`
	Supplier          string  `json:"Supplier" validate:"required,min=1"`
	Notes             *string `json:"Notes"`
	Produk            string  `json:"Produk" validate:"required,min=1"`
	JumlahDus         int     `json:"JumlahDus" validate:"required,min=1"`
	JumlahPcs         int     `json:"JumlahPcs" validate:"required,min=1"`
}

func PengeluranBarangReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams((c)) // Parsing Automate Query Params
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfull to get PengeluranBarang",
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

func PengeluranBarangReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	return c.JSON(fiber.Map{
		"message": "Successfull to get PengeluranBarang detail",
		"status":  200,
	})
}

func PengeluranBarangCreate(c *fiber.Ctx) error {
	var body PengeluranBarangRequest

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
		"message": "Successfull to create PengeluranBarang",
		"status":  200,
		"data":    body,
	})
}

func PengeluranBarangUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	var body PengeluranBarangRequest

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
		"message": "Successfull to update PengeluranBarang",
		"status":  200,
		"data":    body,
	})
}

func PengeluranBarangDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	return c.JSON(fiber.Map{
		"message": "Successfull to delete PengeluranBarang",
		"status":  200,
	})
}
