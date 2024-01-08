package main

import (
	"fmt"
	"net/http"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			// set response code
			w.WriteHeader(400)
			// set response body
			_, err := w.Write([]byte(`{"message': "Hello world!"}`))
			if err != nil {
				fmt.Println("Error writing response body", err)
			}
			return
		}

		// set response code
		w.WriteHeader(200)
		// set response body
		_, err := w.Write([]byte(`pong`))
		if err != nil {
			fmt.Println("Error writing response body", err)
		}

	}

	http.HandleFunc("/ping", handler)

	// run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
