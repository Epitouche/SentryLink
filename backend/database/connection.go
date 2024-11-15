package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "host=postgres user=admin password=password dbname=mydatabase port=5432 sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return conn
}
