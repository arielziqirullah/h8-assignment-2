package database

import (
	"fmt"
	"h8-assignment-2/helpers"
	"h8-assignment-2/models"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup database connection
func SetUpDatabaseConnection() *gorm.DB {

	log.Println("[START] Connecting to database...")

	helpers.LoadEnv()

	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	helpers.LogIfError(err)
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[ERROR] Failed to connect to database!")
	}

	log.Println("[SUCCESS] Connected to database...")

	if os.Getenv("DB_MIGRATE") == "true" {
		log.Println("[START] Migrating database...")
		err := db.AutoMigrate(&models.Orders{}, &models.Items{})
		helpers.LogIfError(err)
		log.Println("[SUCCESS] Migrated database...")
	}

	return db
}

// close database connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	helpers.LogIfError(err)

	dbSQL.Close()
}
