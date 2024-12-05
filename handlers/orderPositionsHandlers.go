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

	err := context.BodyParser(&orderPosition)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&orderPosition).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create order position"})
		return err
	}

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
