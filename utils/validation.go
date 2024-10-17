package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidationQueryParams(c *fiber.Ctx) (string, int, int, string, string, error) {
	// Query Search
	searchQuery := c.Query("search", "")

	// Query Pagination
	pageQuery := c.Query("page", "1")
	sizeQuery := c.Query("size", "5")

	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeQuery)
	if err != nil || size < 1 {
		size = 1
	}

	order := c.Query("order", "asc")
	if order != "asc" && order != "desc" {
		return searchQuery, page, size, order, "", errors.New("invalid order parameter, must be 'asc' or 'desc'")
	}

	// validSortBy := map[string]bool{"id": true, "createdAt": true } // set enum
	sortBy := c.Query("sort_by", "")
	// if sortBy == "" {
	// 	sortBy = "id"
	// }

	return searchQuery, page, size, order, sortBy, nil
}

func ValidationIdParams(c *fiber.Ctx) (int, error) {
	idParams := c.Params("id", "")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		return 0, errors.New("invalid id params, params must be integer")
	}

	return id, nil
}

func ValidateStruct(c *fiber.Ctx, input interface{}) (map[string]string, error) {
	// Memparsing request body ke dalam struct yang diberikan
	if err := c.BodyParser(input); err != nil {
		return nil, errors.New("failed to parse request body")
	}

	// Melakukan validasi struct menggunakan validator
	if err := validate.Struct(input); err != nil {
		// Mengambil kesalahan validasi secara detail
		validationErrors := err.(validator.ValidationErrors)
		errorsMap := make(map[string]string)

		// Mengisi map dengan kesalahan validasi dan pesan error yang dinamis
		for _, err := range validationErrors {
			// Dinamis berdasarkan field dan jenis error
			message := generateErrorMessage(err)
			errorsMap[err.Field()] = message
		}

		// Kembalikan map kesalahan dan error untuk ditangani di controller
		return errorsMap, errors.New("validation failed")
	}

	// Jika validasi berhasil, tidak ada error
	return nil, nil
}

// Fungsi untuk membuat pesan error yang dinamis berdasarkan tag validasi
func generateErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
