package storage

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"github.com/ikardb/BazyDanych/routes"
	"github.com/joho/godotenv"
)

func Setup() {
	err := godotenv.Load("storage/.env")
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := NewConnection(config)
	if err != nil {
		log.Fatal("could not load database")
	}
	err = models.MigrateAll(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	app := fiber.New()
	routes.SetupRoutes(app, db)
	app.Listen(":8080")
}
