package postgreconfigpackage

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDBURI() (string, error) {
	var (
		dbUsername string
		dbName     string
		dbPassword string
		dbHost     string
		dbPort     string
	)

	if err := godotenv.Load(); err != nil {
		return "", err
	}

	dbUsername = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbPassword, dbName), nil
}
