package example

// This file was generated via test-generator

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var h http.Handler = NewRouter()


func TestGetAllTodos(t *testing.T) {

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal("Creating 'GET /todos' request failed!")
	}

	h.ServeHTTP(respRec, req)

	if respRec.Code != 200 {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", 200)
	}
}

