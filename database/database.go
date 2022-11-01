package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {

	// Cargar archivo .env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading database info")
		return nil
	}

	//Cargar variables sql server
	server := os.Getenv("SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DATABASE")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err == nil {
		return db
	}

	return nil

}
