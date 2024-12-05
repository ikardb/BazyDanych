package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type DiningMenuHandler struct {
	DB *gorm.DB
}

func (h *DiningMenuHandler) CreateDiningMenu(context *fiber.Ctx) error {
	diningMenu := models.DiningMenus{}

	err := context.BodyParser(&diningMenu)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&diningMenu).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create dining menu"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "dining menu has been added"})
	return nil
}

func (h *DiningMenuHandler) GetDiningMenus(context *fiber.Ctx) error {
	diningMenus := &[]models.DiningMenus{}

	err := h.DB.Find(diningMenus).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get dining menus"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "dining menus loaded successfully", "data": diningMenus})
	return nil
}

func (h *DiningMenuHandler) GetDiningMenuById(context *fiber.Ctx) error {
	diningMenu := &models.DiningMenus{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_jadlospisu = ?", id).First(diningMenu).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "could find dining menu"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "dining menu fetched succesfully"})
	return nil
}
