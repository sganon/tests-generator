//go:generate tests-generator --specs-file ./specs.yaml --pkg example --handler-func NewRouter
package example

// This file was generated via test-generator

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/json"
)

var h http.Handler = NewRouter()


func TestGetAllTodos1(t *testing.T) {
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

	assert.Equal(t, respRec.Code, 200, "GetAllTodos1: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(`[{ "name": "My First todo", "isFinished": false }]`))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"GetAllTodos1: response body is not matching")
}

func TestAddTodo1(t *testing.T) {
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

	assert.Equal(t, respRec.Code, 201, "AddTodo1: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(`{ "name": "second todo", "isFinished": false }`))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"AddTodo1: response body is not matching")
}

func TestGetAllTodos2(t *testing.T) {
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

	assert.Equal(t, respRec.Code, 200, "GetAllTodos2: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(`[{ "name": "My First todo", "isFinished": false }, { "name": "second todo", "isFinished": false }]`))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"GetAllTodos2: response body is not matching")
}

func TestGetFirstTodo(t *testing.T) {
	var reqBody io.Reader
	if `` != "" {
		reqBody =  bytes.NewBuffer([]byte(``))
	}

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/todos/0", reqBody)
	if err != nil {
		t.Fatal("Creating 'GET /todos/0' request failed!")
	}

	h.ServeHTTP(respRec, req)

	assert.Equal(t, respRec.Code, 200, "GetFirstTodo: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(`{ "name": "My First todo", "isFinished": false }`))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"GetFirstTodo: response body is not matching")
}

func TestGetUnknownTodo(t *testing.T) {
	var reqBody io.Reader
	if `` != "" {
		reqBody =  bytes.NewBuffer([]byte(``))
	}

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/todos/42", reqBody)
	if err != nil {
		t.Fatal("Creating 'GET /todos/42' request failed!")
	}

	h.ServeHTTP(respRec, req)

	assert.Equal(t, respRec.Code, 404, "GetUnknownTodo: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(``))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"GetUnknownTodo: response body is not matching")
}

func TestGetOneTodoWrongParams(t *testing.T) {
	var reqBody io.Reader
	if `` != "" {
		reqBody =  bytes.NewBuffer([]byte(``))
	}

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/todos/fdsd", reqBody)
	if err != nil {
		t.Fatal("Creating 'GET /todos/fdsd' request failed!")
	}

	h.ServeHTTP(respRec, req)

	assert.Equal(t, respRec.Code, 400, "GetOneTodoWrongParams: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(``))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"GetOneTodoWrongParams: response body is not matching")
}

