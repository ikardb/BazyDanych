package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(context *fiber.Ctx) error {
	user := models.Users{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "user has been added"})
	return nil
}

func (h *UserHandler) GetUsers(context *fiber.Ctx) error {
	users := &[]models.Users{}

	err := h.DB.Find(users).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "users loaded successfully", "data": users})
	return nil
}

func (h *UserHandler) GetUserById(context *fiber.Ctx) error {
	user := &models.Users{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_uzytkownika = ?", id).First(user).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "User ID fetched successfully", "data": user})
	return nil
}
