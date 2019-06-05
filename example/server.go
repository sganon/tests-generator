package example

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter returns an http.handler serving the api routes.
// In order to generate test for your routes the handler needs to be exposed
func NewRouter() http.Handler {
	router := httprouter.New()

	type todo struct {
		Name       string `json:"name"`
		IsFinished bool   `json:"isFinished"`
	}
	todos := []todo{
		todo{Name: "My First todo", IsFinished: false},
	}

	returnTodos := func(w http.ResponseWriter) {
		b, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(b)
	}

	router.GET("/todos", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		returnTodos(w)
	})

	router.POST("/todos", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		td := todo{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&td)
		todos = append(todos, td)
		w.WriteHeader(201)
		returnTodos(w)
	})

	return router
}
