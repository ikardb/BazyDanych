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

	sale := models.Sales{}
	err = h.DB.First(&sale, "id_sprzedazy = ?", salePosition.Id_sprzedazy).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "sale not found"})
		return err
	}

	stockLevels := models.StockLevels{}
	err = h.DB.First(&stockLevels, "id_sklepu = ? AND id_produktu = ?", sale.Id_sklepu, salePosition.Id_produktu).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "product not found in stock"})
		return err
	}

	if stockLevels.Ilosc < salePosition.Ilosc {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "not enough product in stock"})
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

	err = h.DB.Model(&stockLevels).Update("ilosc", stockLevels.Ilosc-salePosition.Ilosc).Error
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "could not update stock"})
		return err
	}

	positionCost := float64(*salePosition.Cena_jednostkowa) * float64(salePosition.Ilosc)

	updateQuery := `
        UPDATE sprzedaz
        SET kwota_transakcji = kwota_transakcji + ?
        WHERE Id_sprzedazy = ?
	`
	err = h.DB.Exec(updateQuery, positionCost, salePosition.Id_sprzedazy).Error
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to update order cost",
		})
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
