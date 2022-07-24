package data

import (
	"encoding/json"
	"io"
	"time"
)

// define Product structure (specifying json naming)
type Product struct {
	ID 			int `json:"id"`
	Name 		string `json:"name"`
	Description string `json:"description"`
	Price 		float32 `json:"price"`
	SKU 		string `json:"sku"`
	CreatedOn 	string `json:"-"`
	UpdatedOn 	string `json:"-"`
	DeletedOn 	string `json:"-"`
}

// type Products, list of Product
type Products []*Product

// ToJson function
func (p *Products) ToJson(w io.Writer) error {
	// instantiate new encoder with the receiver io writer
	encoder := json.NewEncoder(w)
	// encode products and possibly return error
	return encoder.Encode(p)
}

// function that returns the product list
func GetProducts() Products {
	return productList
}

// create product list
var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee",
		Price: 0.99,
		SKU: "abcd32",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}