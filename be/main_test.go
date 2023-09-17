package main

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	handler := helloHandler()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)
}
