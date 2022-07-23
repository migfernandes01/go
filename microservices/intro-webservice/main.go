package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// execute function when we got incoming request to specified path
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		// read everything in the request body (returns data and error)
		d, err := ioutil.ReadAll(r.Body)

		// if no error is returned from reading req body
		if(err != nil) {
			// throw an http error
			http.Error(rw, "Error message", http.StatusBadRequest)
			return
		}

		log.Printf("Request body data: %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	// serve on port 9090
	http.ListenAndServe(":9090", nil)
}