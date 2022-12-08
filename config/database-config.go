package config

import (
	"backend/entity"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")

	}
	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbHost := os.Getenv("PGHOST")
	dbName := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=7634 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.UserPromo{}, &entity.Address{}, &entity.Shipping{}, &entity.Payment{}, &entity.AddOn{}, &entity.AddOnShipping{}, &entity.Category{}, &entity.Size{}, &entity.Promo{}, &entity.User{}, &entity.Transaction{}, &entity.SourceOfFund{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPostgres, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbPostgres.Close()
}
