package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateOrder_WithMock(t *testing.T) {
	// stworzenie mocka bazy
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// podlaczenie mocka do GORMa
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	handler := &OrderHandler{DB: gormDB}

	app := fiber.New()
	app.Post("/orders", handler.CreateOrder)

	// przykladowe dane wejsciowe
	orderData := map[string]interface{}{
		"id_sklepu":      1,
		"id_uzytkownika": 2,
	}
	body, _ := json.Marshal(orderData)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "zamowienie"`).
		WithArgs(
			int16(1),
			int16(2),
			sqlmock.AnyArg(),
			0.0,
			false,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id_zamowienia"}).AddRow(1))
	mock.ExpectCommit()

	req := httptest.NewRequest("POST", "/orders", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, mock.ExpectationsWereMet())
}
