package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MungaiVic/inventory/pkg/config"
	"github.com/MungaiVic/inventory/pkg/models"
	"github.com/MungaiVic/inventory/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"github.com/mgutz/ansi"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeZone:   "Africa/Nairobi",
		TimeFormat: "2006-01-02 15:04:05",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1").(*fiber.Group)
	routes.SetupItemRoutes(v1, db)
	routes.SetupUserRoutes(v1, db)
}

func initDatabase() *gorm.DB {
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
		redify := ansi.ColorFunc("red")
		msg := redify(fmt.Sprintf("%s", err))
		fmt.Println(msg)
		log.Fatal("Could not migrate db on Items")
	}
	err = models.MigrateUsers(db)
	if err != nil {
		redify := ansi.ColorFunc("red")
		msg := redify(fmt.Sprintf("%s", err))
		fmt.Println(msg)
		log.Fatal("Could not migrate db on Users")
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
	log.Fatal(app.Listen(":5000"))

}
