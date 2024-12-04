package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type OrderHandler struct {
	DB *gorm.DB
}

func (h *OrderHandler) CreateOrder(context *fiber.Ctx) error {
	order := models.Orders{}
	order.Data_zamowienia = time.Now()

	err := context.BodyParser(&order)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&order).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create order"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order has been added"})
	return nil
}

func (h *OrderHandler) GetOrders(context *fiber.Ctx) error {
	orders := &[]models.Orders{}

	err := h.DB.Find(orders).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get orders"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "orders loaded successfully", "data": orders})
	return nil
}

func (h *OrderHandler) GetOrderById(context *fiber.Ctx) error {
	order := &models.Orders{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_zamowienia = ?", id).First(order).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "order not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order ID fetched successfully", "data": order})
	return nil
}
