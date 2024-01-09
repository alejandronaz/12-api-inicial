package main

import (
	"encoding/json"
	"fmt"
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
		reader := r.Body // r.Body returns io.ReadCloser which implements the interface io.Reader

		// OPCION 1: UTILIZAR UNMARSHAL

		// bytesMessage, err := io.ReadAll(reader)
		// if err == io.EOF {
		// 	// set response code
		// 	w.WriteHeader(400)

		// 	_, err := w.Write([]byte(`error while reading body`))
		// 	if err != nil {
		// 		fmt.Println("Error writing response body while parsing it: ", err)
		// 	}
		// 	return
		// }

		// // decode body bytes to struct
		// var body Greeting
		// if err := json.Unmarshal(bytesMessage, &body); err != nil {
		// 	fmt.Println("Error parsing body: ", err) // for example, if body is empty or its malformed

		// 	// set response code
		// 	w.WriteHeader(400)
		// 	_, err := w.Write([]byte(`body has not the correct format`))
		// 	if err != nil {
		// 		fmt.Println("Error writing response body while parsing it: ", err)
		// 	}
		// 	return
		// }

		// OPCION 2: UTILIZAR DECODER

		// creo un decoder
		dc := json.NewDecoder(reader)

		// decode json to struct
		var body Greeting
		err := dc.Decode(&body)
		if err != nil {
			fmt.Println(err)
			return
		}

		// set response code
		w.WriteHeader(200)

		// set response body
		resMessage := fmt.Sprintf("Hello %s %s", body.FirstName, body.LastName)
		_, err = w.Write([]byte(resMessage))
		if err != nil {
			fmt.Println("Error writing response body: ", err)
			return
		}

	}

	http.HandleFunc("/greetings", handler)

	// run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
