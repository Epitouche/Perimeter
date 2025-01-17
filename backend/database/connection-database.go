package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection establishes a connection to the PostgreSQL database using environment variables for configuration.
// It retrieves the following environment variables:
// - DB_HOST: the database host
// - DB_PORT: the database port
// - POSTGRES_USER: the database user
// - POSTGRES_PASSWORD: the database password
// - POSTGRES_DB: the database name
// If any of these environment variables are not set, the function will panic with an appropriate error message.
// The function returns a pointer to a gorm.DB instance representing the database connection.
// If the connection fails, the function will panic with a "failed to connect database" message.
func Connection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		panic("DB_HOST is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		panic("DB_PORT is not set")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		panic("POSTGRES_USER is not set")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		panic("POSTGRES_PASSWORD is not set")
	}

	dbname := os.Getenv("POSTGRES_DB")
	if dbname == "" {
		panic("POSTGRES_DB is not set")
	}

	dsn := "host=" + host + " user=" + user + " password=" + password +
		" dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("Connection to database established")

	// if os.Getenv("GIN_MODE") != "release" {
	// 	conn = conn.Debug() // Enable debugging
	// }
	return conn
}
