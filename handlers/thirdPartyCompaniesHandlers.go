package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type ThirdPartyCompanyHandler struct {
	DB *gorm.DB
}

func (h *ThirdPartyCompanyHandler) CreateThirdPartyCompany(context *fiber.Ctx) error {
	thirdPartyCompany := models.ThirdPartyCompanies{}

	err := context.BodyParser(&thirdPartyCompany)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&thirdPartyCompany).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create third party company"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "third party company has been added"})
	return nil
}

func (h *ThirdPartyCompanyHandler) GetThirdPartyCompanies(context *fiber.Ctx) error {
	thirdPartyCompanies := &[]models.ThirdPartyCompanies{}

	err := h.DB.Find(thirdPartyCompanies).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get third party companies"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "third party companies loaded successfully", "data": thirdPartyCompanies})
	return nil
}

func (h *ThirdPartyCompanyHandler) GetThirdPartyCompanyById(context *fiber.Ctx) error {
	thirdPartyCompany := &models.ThirdPartyCompanies{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_firmy = ?", id).First(thirdPartyCompany).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "third party company not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "third party company ID fetched successfully", "data": thirdPartyCompany})
	return nil
}

func (h *ThirdPartyCompanyHandler) GetThirdPartyCompanyOfferPositionsById(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	thirdPartyCompanyPositions := []models.ThirdPartyCompanyPositions{}
	query := `
		SELECT *
		FROM oferta_firmy_zewnetrznej p
		JOIN firma_zewnetrzna z ON z.id_firmy = p.id_firmy
		WHERE z.id_firmy = ?
		`
	err := h.DB.Raw(query, id).Scan(&thirdPartyCompanyPositions).Error

	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "third party company offer positions not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "third party offer positions fetched successfully", "data": thirdPartyCompanyPositions})
	return nil
}
