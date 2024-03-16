package config

import (
	"fmt"
	"os"

	"github.com/luthfirahman/user-svc/internal/utils/constant"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	if os.Getenv("APP_ENV") != constant.ENUM_RUN_PRODUCTION {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbTimezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v TimeZone=%v", dbHost, dbPort, dbUser, dbPass, dbName, dbTimezone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
