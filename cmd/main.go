package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MungaiVic/inventory/pkg/config"
	"github.com/MungaiVic/inventory/pkg/models"
	"github.com/MungaiVic/inventory/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupRoutes(app *fiber.App, db *gorm.DB) {
    routes.SetupItemRoutes(app, db)
}


func initDatabase() *gorm.DB{
	// load in connection configuration for DB
	configuration := &config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// Create new connection using configurations
	db, err := config.NewConnection(configuration)

	if err != nil {
		log.Fatal("Could not load the database")
	}
	// Migrate models
	err = models.MigrateItems(db)
	if err != nil {
		log.Fatal("Could not migrate db")
	}
	fmt.Println("DB migrated!")
	return db

}

func main() {
	// Load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize app and routes
	app := fiber.New()
	db = initDatabase()
	setupRoutes(app, db)
	app.Listen(":5000")
	
}
