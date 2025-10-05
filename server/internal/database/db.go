package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"apiGo/internal/models"
)

var DB *gorm.DB

func Setup() {
	dsn := generateDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Conexión a la base de datos fallida: %v", err)
	}

	if err := migrateModels(); err != nil {
		log.Fatalf("")
	}
}

func generateDSN() string {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("Por favor, configure una base de datos")
	}
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "require" // default SSL mode
	}
	timeZone := os.Getenv("DB_TIMEZONE")
	if timeZone == "" {
		timeZone = "UTC" // default timezone
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode, timeZone,
	)
}

func migrateModels() error {
	models := []any{
		&models.Departamento{},
		&models.Ciudad{},
		&models.Barrio{},
	}

	return DB.AutoMigrate(models...)
}
