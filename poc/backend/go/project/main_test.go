package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter() // Use the setupRouter function from main.go

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
}

func TestWordRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"hello"}`, w.Body.String())
}
