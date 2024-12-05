package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type DiningMenuPositionHandler struct {
	DB *gorm.DB
}

func (h *DiningMenuPositionHandler) CreateDiningMenuPosition(context *fiber.Ctx) error {
	position := models.DiningMenuPositions{}

	err := context.BodyParser(&position)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&position).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create position"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "position has been added"})
	return nil
}

func (h *DiningMenuPositionHandler) GetDiningMenuPositions(context *fiber.Ctx) error {
	positions := &[]models.DiningMenuPositions{}

	err := h.DB.Find(positions).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get positions"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "positions loaded successfully", "data": positions})
	return nil
}
