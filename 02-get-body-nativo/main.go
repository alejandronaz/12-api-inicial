package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Greeting struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {

		// check if method is POST
		if r.Method != http.MethodPost {
			// set response code
			w.WriteHeader(400)
			// set response body
			_, err := w.Write([]byte(`Only POST method is allowed`))
			if err != nil {
				fmt.Println("Error writing response body", err)
			}
			return
		}

		// get body and parse it to bytes
		reader := r.Body

		bytesMessage, err := io.ReadAll(reader)
		if err == io.EOF {
			_, err := w.Write([]byte(`body is missing`))
			if err != nil {
				fmt.Println("Error writing response body while parsing it", err)
			}
			return
		}

		// parse bytes to struct
		var body Greeting
		if err := json.Unmarshal([]byte(bytesMessage), &body); err != nil {
			fmt.Println("Error parsing body", err)
			return
		}

		// set response code
		w.WriteHeader(200)
		// set response body
		resMessage := fmt.Sprintf("Hello %s %s", body.FirstName, body.LastName)
		_, err = w.Write([]byte(resMessage))
		if err != nil {
			fmt.Println("Error writing response body", err)
			return
		}

	}

	http.HandleFunc("/greetings", handler)

	// run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
