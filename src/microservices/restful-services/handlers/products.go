package handlers

import (
	"log"
	"main/src/microservices/restful-services/data"
	"net/http"
	"regexp"
	"strconv"
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
	// if its a POST req
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	// if its a PUT req
	if r.Method == http.MethodPut {
		// expect ID in url - extract it
		regex := regexp.MustCompile(`/([0-9]+)`)
		stringGroup := regex.FindAllStringSubmatch(r.URL.Path, -1)
		// if stringGroup that matched that regex is not 1, http error
		if len(stringGroup) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		// if stringGroup[0] that matched that regex is not 2, http error
		if len(stringGroup[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		// get id from URL and convert to int
		idString := stringGroup[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Unable to convert string to integer", http.StatusBadRequest)
		}

		// call updateProducts passing the id, rw and r
		p.updateProduct(id, rw, r)
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

// add product
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST req")
	// allocate space in memory for a new "Product"
	product := &data.Product{}
	// call FromJson passing the req body
	err := product.FromJson(r.Body)
	// throw error if there was an error
	if err != nil {
		http.Error(rw, "Unable to convert from JSON", http.StatusBadRequest)
	}

	p.log.Printf("New Product: %#v", product)

	// Add product
	data.AddProduct(product)
}

//
func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	// allocate space in memory for a new "Product"
	product := &data.Product{}
	// call FromJson passing the req body
	err := product.FromJson(r.Body)
	// throw error if there was an error
	if err != nil {
		http.Error(rw, "Unable to convert from JSON", http.StatusBadRequest)
	}

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