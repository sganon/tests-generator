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

	router.GET("/todos", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		todos := []todo{
			todo{Name: "My First todo", IsFinished: false},
		}
		b, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(b)
	})
	return router
}
