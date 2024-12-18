package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tom-Mendy/SentryLink/test"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	postgresContainer, err := test.CreatePostgresContainer(ctx)
	assert.NoError(t, err, "failed to create Postgres container")
	assert.NotNil(t, postgresContainer, "failed to create Postgres container")

	defer func() {
		err := postgresContainer.Terminate(ctx)
		assert.NoError(t, err)
	}()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String(), "unexpected response body")
}

func TestWordRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"hello"}`, w.Body.String())
}
