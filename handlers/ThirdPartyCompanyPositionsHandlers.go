package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type ThirdPartyCompanyPositionHandler struct {
	DB *gorm.DB
}

func (h *ThirdPartyCompanyPositionHandler) CreateThirdPartyCompanyPosition(context *fiber.Ctx) error {
	thirdPartyCompanyPosition := models.ThirdPartyCompanyPositions{}

	err := context.BodyParser(&thirdPartyCompanyPosition)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&thirdPartyCompanyPosition).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create third party company position"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "third party company position has been added"})
	return nil
}

func (h *ThirdPartyCompanyPositionHandler) GetThirdPartyCompanyPositions(context *fiber.Ctx) error {
	thirdPartyCompanyPositions := &[]models.ThirdPartyCompanyPositions{}

	err := h.DB.Find(thirdPartyCompanyPositions).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get third party company positions"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "third party company positions loaded successfully", "data": thirdPartyCompanyPositions})
	return nil
}
