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
	order.Koszt_zamowienia = 0
	order.Wczytane_do_stanu = false

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

func (h *OrderHandler) GetOrderPositions(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	orderPositions := []models.OrderPositionsWithProductName{}
	query := `
		SELECT 
			p.id_pozycji_zamowienia,
			p.id_zamowienia,
			p.ilosc_produktu,
			produkt.nazwa,
			produkt.cena
		FROM pozycja_zamowienia p
		JOIN zamowienie z ON z.id_zamowienia = p.id_zamowienia
		JOIN pozycja_jadlospisu ON pozycja_jadlospisu.id_pozycji_jadlospisu = p.id_pozycji_jadlospisu
		JOIN produkt ON produkt.id_produktu = pozycja_jadlospisu.id_produktu
		WHERE z.id_zamowienia = ?
	`

	err := h.DB.Raw(query, id).Scan(&orderPositions).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "order positions not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order positions fetched successfully", "data": orderPositions})
	return nil
}

func (h *OrderHandler) MigrateToStock(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Order ID is required"})
		return nil
	}

	order := models.Orders{}
	err := h.DB.Where("id_zamowienia = ?", id).First(&order).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Order not found"})
	}

	if order.Wczytane_do_stanu {
		context.Status(http.StatusConflict).JSON(&fiber.Map{"message": "Order has already been migrated to stock"})
		return nil
	}

	orderPositions := []models.OrderPositions{}

	err = h.DB.Where("id_zamowienia = ?", id).Find(&orderPositions).Error
	if err != nil || len(orderPositions) == 0 {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Order positions not found"})
		return err
	}

	for _, orderPosition := range orderPositions {
		diningMenuPosition := models.DiningMenuPositions{}
		err = h.DB.Where("id_pozycji_jadlospisu = ?", orderPosition.Id_pozycji_jadlospisu).First(&diningMenuPosition).Error
		if err != nil {
			continue
		}

		product := models.Products{}
		err = h.DB.Where("id_produktu = ?", diningMenuPosition.Id_produktu).First(&product).Error
		if err != nil {
			continue
		}

		stockLevel := models.StockLevels{}
		err = h.DB.Where("id_sklepu = ? AND id_produktu = ?", order.Id_sklepu, product.Id_produktu).First(&stockLevel).Error

		if err != nil {
			newStockLevel := models.StockLevels{
				Id_sklepu:   order.Id_sklepu,
				Id_produktu: product.Id_produktu,
				Ilosc:       orderPosition.Ilosc_produktu,
			}
			err = h.DB.Create(&newStockLevel).Error

			if err != nil {
				context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed to create stock level"})
				return err
			}
		} else {
			stockLevel.Ilosc += orderPosition.Ilosc_produktu
			err = h.DB.Save(&stockLevel).Error

			if err != nil {
				context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed to update stock level"})
				return err
			}
		}
	}

	err = h.DB.Model(&order).Update("Wczytane_do_stanu", true).Error
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed to update order flag"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Products migrated to stock levels successfully"})
	return nil
}

func (h *OrderHandler) GetOrdersByShopId(context *fiber.Ctx) error {
	orders := &[]models.Orders{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_sklepu = ?", id).Find(orders).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "order not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order ID fetched successfully", "data": orders})
	return nil
}
