package handlers

import (
	"context"
	"log"
	"microservices/restful-services/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	log *log.Logger
}

// take logger returns Products handler
func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

// get products and return them in JSON obj
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// call products getter from data package
	lp := data.GetProducts()

	// call ToJSON method on Products passing response writer (type of io writer)
	err :=  lp.ToJson(rw)

	// throw error if there was an error
	if err != nil {
		http.Error(rw, "Unable to convert to JSON", http.StatusInternalServerError)
	}
}

// add product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST req")

	// get produt from request context and cast it
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.log.Printf("New Product: %#v", product)

	// Add product
	data.AddProduct(product)
}

//
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// get url params passing request to mux.Vars
	vars := mux.Vars(r) 
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.log.Println("Handle PUT req", id)
	// get produt from request context and cast it
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	// call updateProduct passing the id and product getting error back
	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}

type KeyProduct struct{}

// middleware to validate product
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// allocate space in memory for a new "Product"
		product := &data.Product{}
		// call FromJson passing the req body
		err := product.FromJson(r.Body)
		// throw error if there was an error
		if err != nil {
			http.Error(rw, "Unable to convert from JSON", http.StatusBadRequest)
			return
		}

		// add valid product to context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		// call next passing req with context
		next.ServeHTTP(rw, r)
	})
}