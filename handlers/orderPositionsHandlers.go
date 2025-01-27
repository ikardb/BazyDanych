package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type OrderPositionHandler struct {
	DB *gorm.DB
}

func (h *OrderPositionHandler) CreateOrderPosition(context *fiber.Ctx) error {
	orderPosition := models.OrderPositions{}
	product := models.Products{}
	diningMenuPosition := models.DiningMenuPositions{}

	err := context.BodyParser(&orderPosition)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Where("id_pozycji_jadlospisu = ?", orderPosition.Id_pozycji_jadlospisu).First(&diningMenuPosition).Error
	if err != nil {
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Menu item not found",
		})
	}

	err = h.DB.Where("id_produktu = ?", diningMenuPosition.Id_produktu).First(&product).Error
	if err != nil {
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Menu item not found",
		})
	}

	positionCost := float64(product.Cena) * float64(orderPosition.Ilosc_produktu)

	// Start a new transaction
	tx := h.DB.Begin()

	err = tx.Create(&orderPosition).Error
	if err != nil {
		tx.Rollback()
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create order position"})
		return err
	}

	updateQuery := `
		UPDATE zamowienie
		SET koszt_zamowienia = koszt_zamowienia + ?
		WHERE id_zamowienia = ?
	`
	err = tx.Exec(updateQuery, positionCost, orderPosition.Id_zamowienia).Error
	if err != nil {
		tx.Rollback()
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to update order cost",
		})
	}

	// Commit the transaction
	tx.Commit()

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order position has been added"})
	return nil
}

func (h *OrderPositionHandler) GetOrderPositions(context *fiber.Ctx) error {
	orderPositions := &[]models.OrderPositions{}

	err := h.DB.Find(orderPositions).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get order positions"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order positions loaded successfully", "data": orderPositions})
	return nil
}
