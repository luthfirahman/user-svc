package config

import (
	"fmt"
	"log"
	"os"

	"github.com/luthfirahman/user-svc/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {

	utils.LoadEnv()
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v TimeZone=Asia/Jakarta", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
