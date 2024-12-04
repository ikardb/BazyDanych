package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type KitchenHandler struct {
	DB *gorm.DB
}

func (h *KitchenHandler) CreateKitchen(context *fiber.Ctx) error {
	kitchen := models.Kitchens{}

	err := context.BodyParser(&kitchen)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&kitchen).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create kitchen"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "kitchen has been added"})
	return nil
}

func (h *KitchenHandler) GetKitchens(context *fiber.Ctx) error {
	kitchens := &[]models.Kitchens{}

	err := h.DB.Find(kitchens).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get kitchens"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "kitchens loaded successfully", "data": kitchens})
	return nil
}

func (h *KitchenHandler) GetKitchenById(context *fiber.Ctx) error {
	kitchen := &models.Kitchens{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_kuchni = ?", id).First(kitchen).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "kitchen not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "kitchen ID fetched successfully", "data": kitchen})
	return nil
}
