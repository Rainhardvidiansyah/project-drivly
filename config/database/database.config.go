package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseConfig struct {
	DatabaseName     string
	DatabasePassword string
	DatabaseType     string
	DatabasePort     string
	DatabaseDriver   string
}

// CALL ENVIRONMENT
func LoadConfig() *DatabaseConfig { //return *DatabaseConfig, ini berarti LoadConfig menjadi DatabaseConfig
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print("😪😪😯😯 env cannot be loaded 😪😪😯😯")
		panic(err)
	}

	return &DatabaseConfig{
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseType:     os.Getenv("DATABASE_TYPE"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseDriver:   os.Getenv("DATABASE_DRIVER"),
	}

}

// CONNECT TO DATABASE -- CALL THIS FUNCTION AT MAIN FUNCTION
func InitializeDatabase(cfg *DatabaseConfig) {

	dsn := fmt.Sprintf(
		"host=localhost port=%s user=postgres password=%s dbname=%s sslmode=disable",
		cfg.DatabasePort, cfg.DatabasePassword, cfg.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	log.Println(" Database connected successfully! ")
}
