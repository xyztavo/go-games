package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConnString() string {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	connStr := os.Getenv("PG_CONNECTION_STRING")

	return connStr
}

func GetPort() string {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	return port
}
