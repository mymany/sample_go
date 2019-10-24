package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
		router := setUpRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
}