package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_STRING")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.Exec("SHOW TABLES")
	if err != nil {
		log.Fatalf("Failed to Connect to Database")
	}
}
