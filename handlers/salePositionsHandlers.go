package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type SalePositionHandler struct {
	DB *gorm.DB
}

func (h *SalePositionHandler) CreateSalePosition(context *fiber.Ctx) error {
	salePosition := models.SalePositions{}

	err := context.BodyParser(&salePosition)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	if salePosition.Cena_jednostkowa == nil {
		product := models.Products{}
		err := h.DB.First(&product, "id_produktu = ?", salePosition.Id_produktu).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "product not found"})
			return err
		}
		salePosition.Cena_jednostkowa = &product.Cena
	}

	err = h.DB.Create(&salePosition).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create saleposition"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sale position has been added"})
	return nil
}

func (h *SalePositionHandler) GetSalePositions(context *fiber.Ctx) error {
	salePositions := &[]models.SalePositions{}

	err := h.DB.Find(salePositions).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get sale positions"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sale positions loaded successfully", "data": salePositions})
	return nil
}
