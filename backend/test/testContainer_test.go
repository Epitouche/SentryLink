package test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePostgresContainer(t *testing.T) {
	ctx := context.Background()

	// Set environment variables for the test
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpassword")
	os.Setenv("POSTGRES_DB", "testdb")

	container, err := CreatePostgresContainer(ctx)
	assert.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, container, "failed to create Postgres container")

	// Clean up the container after the test
	defer func() {
		err := container.Terminate(ctx)
		assert.NoError(t, err)
	}()

	// Verify the environment variables are set correctly
	// assert.NotEqual(t, "5432", os.Getenv("DB_PORT"))
}
