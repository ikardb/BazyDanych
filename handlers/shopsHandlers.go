package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type ShopHandler struct {
	DB *gorm.DB
}

func (h *ShopHandler) CreateShop(context *fiber.Ctx) error {
	shop := models.Shops{}

	err := context.BodyParser(&shop)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&shop).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create shop"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "shop has been added"})
	return nil
}

func (h *ShopHandler) GetShops(context *fiber.Ctx) error {
	shops := &[]models.Shops{}

	err := h.DB.Find(shops).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get shops"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Shops loaded successfully", "data": shops})
	return nil
}

func (h *ShopHandler) GetShopById(context *fiber.Ctx) error {
	shop := &models.Shops{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_sklepu = ?", id).First(shop).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "shop not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Shop ID fetched successfully", "data": shop})
	return nil
}
