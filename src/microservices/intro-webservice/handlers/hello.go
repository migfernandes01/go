package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// struct that contains log class
type Hello struct {
	log *log.Logger
}

// function that gets a logger and return Hello handler
func NewHello(log *log.Logger) *Hello{
	return &Hello{log}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("Hello World")

	// read everything in the request body (returns data and error)
	d, err := ioutil.ReadAll(r.Body)

	// if no error is returned from reading req body
	if(err != nil) {
		// throw an http error
		http.Error(rw, "Error message", http.StatusBadRequest)
		return
	}

	// print formatted data from req body
	fmt.Fprintf(rw, "Hello %s", d)
}