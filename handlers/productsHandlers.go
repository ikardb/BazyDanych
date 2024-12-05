package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (h *ProductHandler) CreateProduct(context *fiber.Ctx) error {
	product := models.Products{}

	err := context.BodyParser(&product)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&product).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create product"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "product has been added"})
	return nil
}

func (h *ProductHandler) GetProducts(context *fiber.Ctx) error {
	products := &[]models.Products{}

	err := h.DB.Find(products).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get products"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "products loaded successfully", "data": products})
	return nil
}

func (h *ProductHandler) GetProductById(context *fiber.Ctx) error {
	product := &models.Products{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_produktu = ?", id).First(product).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "product not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "product ID fetched successfully", "data": product})
	return nil
}
