package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	// create a new endpoint GET "/hello-world"
	router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		// set response code
		w.WriteHeader(200)
		// set response body
		code, err := w.Write([]byte(`{"message': "Hello world!"}`))
		if err != nil {
			fmt.Println("Error writing response body", err)
		}
		fmt.Println("No se que es el int que retorna -> ", code)
	})

	// run server
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
