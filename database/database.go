package database

import (
	"fmt"
	"log"
	"auth_api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config.LoadEnv()

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_NAME"),
	)

	// Open a connection to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("Database connected successfully!")
}
