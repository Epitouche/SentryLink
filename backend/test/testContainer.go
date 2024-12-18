package test

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreatePostgresContainer(ctx context.Context) (testcontainers.Container, error) {
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

	// Create a new container with PostgreSQL
	dbName := dbname
	dbUser := user
	dbPassword := password

	postgresContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(20*time.Second)),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)

		return postgresContainer, err
	}

	host, err = postgresContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get container host: %s", err)
	}

	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("failed to get container port: %s", err)
	}

	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", mappedPort.Port())

	return postgresContainer, nil
}
