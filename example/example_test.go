package example

// This file was generated via test-generator

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var h http.Handler = NewRouter()


func TestGetAllTodos(t *testing.T) {
	var reqBody io.Reader
	if `` != "" {
		reqBody =  bytes.NewBuffer([]byte(``))
	}

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/todos", reqBody)
	if err != nil {
		t.Fatal("Creating 'GET /todos' request failed!")
	}

	h.ServeHTTP(respRec, req)

	assert.Equal(t, respRec.Code, 200, "GetAllTodos: unexpected response code")

	b, err := ioutil.ReadAll(respRec.Body)
	if err != nil {
		t.Fatal("error reading response body", err)
	}
	body := strings.ReplaceAll(string(b), " ", "")
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")

	ref := strings.ReplaceAll(`[{ "name": "My First todo", "isFinished": false }]`, " ", "")
	ref = strings.ReplaceAll(ref, "\t", "")
	ref = strings.ReplaceAll(ref, "\n", "")

	assert.Equal(t, body, ref, "GetAllTodos: response body is not matching")
}

func TestAddTodo(t *testing.T) {
	var reqBody io.Reader
	if `{"name": "second todo"}` != "" {
		reqBody =  bytes.NewBuffer([]byte(`{"name": "second todo"}`))
	}

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/todos", reqBody)
	if err != nil {
		t.Fatal("Creating 'POST /todos' request failed!")
	}

	h.ServeHTTP(respRec, req)

	assert.Equal(t, respRec.Code, 201, "AddTodo: unexpected response code")

	b, err := ioutil.ReadAll(respRec.Body)
	if err != nil {
		t.Fatal("error reading response body", err)
	}
	body := strings.ReplaceAll(string(b), " ", "")
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")

	ref := strings.ReplaceAll(`[{ "name": "My First todo", "isFinished": false }, { "name": "second todo", "isFinished": false }]`, " ", "")
	ref = strings.ReplaceAll(ref, "\t", "")
	ref = strings.ReplaceAll(ref, "\n", "")

	assert.Equal(t, body, ref, "AddTodo: response body is not matching")
}

