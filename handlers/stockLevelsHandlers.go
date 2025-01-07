package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type StockLevelHandler struct {
	DB *gorm.DB
}

func (h *StockLevelHandler) CreateStockLevel(context *fiber.Ctx) error {
	stockLevel := models.StockLevels{}

	err := context.BodyParser(&stockLevel)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&stockLevel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create stock level"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "stock level has been added"})
	return nil
}

func (h *StockLevelHandler) GetStockLevels(context *fiber.Ctx) error {
	stockLevels := &[]models.StockLevels{}

	err := h.DB.Find(stockLevels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get stock levels"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "stock levels loaded successfully", "data": stockLevels})
	return nil
}

func (h *StockLevelHandler) GetStockLevelById(context *fiber.Ctx) error {
	stockLevel := &models.StockLevels{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_stanu = ?", id).First(stockLevel).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "stock level not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "stock level ID fetched successfully", "data": stockLevel})
	return nil
}

func (h *StockLevelHandler) GetStockLevelByShopId(context *fiber.Ctx) error {
	stockLevels := &[]models.StockLevels{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_sklepu = ?", id).Find(stockLevels).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "stock level not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "stock level ID fetched successfully", "data": stockLevels})
	return nil
}
