package database

import (
	"Backend/config"
	"Backend/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var (
	DB *gorm.DB
)

// ConnectDB connects to the database
func ConnectDB() {
	var err error

	// Parse the database port from the configuration
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	// Create the Data Source Name (DSN) for connecting to the PostgreSQL database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// Open a connection to the database using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Auto-migrate the necessary models
	DB.AutoMigrate(&model.user{})

	fmt.Println("Connection opened to the database")
	fmt.Println("Database migrated")
}
