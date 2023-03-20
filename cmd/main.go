package main

import (
	"fmt"
	"inv-v2/internal/config"
	"inv-v2/internal/handlers"
	"inv-v2/internal/models"
	"inv-v2/internal/repository"
	"inv-v2/internal/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load(".env")

	if enverr != nil {
		log.Fatal(enverr)
	}
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeZone:   "Africa/Nairobi",
		TimeFormat: "2006-01-02 15:04:05",
	}))
	v1 := app.Group("/api/v1").(*fiber.Group)
	// Set up the database
	pgConfigs := config.PostgresConfig{
		Host:     os.Getenv("PG_DB_HOST"),
		Port:     os.Getenv("PG_DB_PORT"),
		Password: os.Getenv("PG_DB_PASS"),
		User:     os.Getenv("PG_DB_USER"),
		DBName:   os.Getenv("PG_DB_NAME"),
		SSLMode:  os.Getenv("PG_DB_SSLMODE"),
	}
	// Get new connection
	dbConn, err := config.NewPostgresConnection(pgConfigs)
	if err != nil {
		panic(err)
	}
	// Migrate the Item model
	err = dbConn.AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	}
	// Migrate the User model
	err = dbConn.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database successfully migrated!")
	itemDAO := repository.NewItemConnection(dbConn)
	userDAO := repository.NewUserConnection(dbConn)
	itemSVC := service.NewItemService(itemDAO)
	userSVC := service.NewUserService(userDAO)
	authSVC := service.NewAuthService(userDAO)

	// Set up the routes
	handlers.SetupItemRoutes(v1, itemSVC)
	handlers.SetupUserRoutes(v1, userSVC)
	handlers.SetupAuthRoutes(v1, authSVC)
	// Start the server
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))

}
