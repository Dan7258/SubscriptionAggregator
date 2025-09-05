package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"subagg/internal/config"
	"subagg/internal/handlers"
	"subagg/internal/models"
	"subagg/internal/routes"
)

func main() {
	config.Init()
	gorm, err := ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db := models.NewDatabase(gorm)
	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewHandlers(db)

	router := gin.Default()

	routes.RegisterRoutes(router, h)

	router.Run(":8080")
}

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"))
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
