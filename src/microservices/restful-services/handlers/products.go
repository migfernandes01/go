package handlers

import (
	"log"
	"main/src/microservices/restful-services/data"
	"net/http"
)

type Products struct {
	log *log.Logger
}

// take logger returns Products handler
func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

// HANDLER 
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// if it's a GET req
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusNotImplemented)
}

// get products and return them in JSON obj
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	// call products getter from data package
	lp := data.GetProducts()

	// call ToJSON method on Products passing response writer (type of io writer)
	err :=  lp.ToJson(rw)

	// throw error if there was an error
	if err != nil {
		http.Error(rw, "Unable to convert to JSON", http.StatusInternalServerError)
	}
}