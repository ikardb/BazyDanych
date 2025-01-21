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
	id := context.Params("id")

	searchQuery := context.Query("q", "")

	if id == "" {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"message": "Shop ID cannot be empty"})
	}

	stockLevels := &[]models.StockLevelWithProduct{}

	query := `
	SELECT 
		s.id_stanu,
		s.id_produktu,
		s.ilosc,
		s.id_sklepu,
		produkt.nazwa
	FROM stan_magazynowy s
	JOIN produkt ON produkt.id_produktu = s.id_produktu
	WHERE s.id_sklepu = ?
	`

	if searchQuery != "" {
		query += ` AND produkt.nazwa ILIKE ?`
	}

	var err error
	if searchQuery != "" {
		err = h.DB.Raw(query, id, searchQuery+"%").Scan(&stockLevels).Error
	} else {
		err = h.DB.Raw(query, id).Scan(&stockLevels).Error
	}

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{"message": "Could not fetch stock levels for the shop"})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{"message": "Stock levels loaded successfully", "data": stockLevels})
}
