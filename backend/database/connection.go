package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		panic("DB_HOST is not set")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		panic("DB_USER is not set")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		panic("DB_PASSWORD is not set")
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		panic("DB_NAME is not set")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		panic("DB_PORT is not set")
	}

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("Connection to database established")

	os.Getenv("GIN_MODE") // Set
	if os.Getenv("GIN_MODE") != "release" {
		conn = conn.Debug() // Enable debugging
	}
	return conn
}
