package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type SaleHandler struct {
	DB *gorm.DB
}

func (h *SaleHandler) CreateSale(context *fiber.Ctx) error {
	sale := models.Sales{}
	sale.Data_sprzedazy = time.Now()

	err := context.BodyParser(&sale)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&sale).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create sale"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sale has been added"})
	return nil
}

func (h *SaleHandler) GetSales(context *fiber.Ctx) error {
	sales := &[]models.Sales{}

	err := h.DB.Find(sales).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get sales"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sales loaded successfully", "data": sales})
	return nil
}

func (h *SaleHandler) GetSaleById(context *fiber.Ctx) error {
	sale := &models.Sales{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_sprzedazy = ?", id).First(sale).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "sale not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sale ID fetched successfully", "data": sale})
	return nil
}
