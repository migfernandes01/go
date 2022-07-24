package data

import (
	"encoding/json"
	"fmt"
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

// ToJson function (uses Products, receives ioWriter, returns error)
func (p *Products) ToJson(w io.Writer) error {
	// instantiate new encoder with the received io writer
	encoder := json.NewEncoder(w)
	// encode products and possibly return error
	return encoder.Encode(p)
}

// decode JSON and write to products (uses Product, receives ioReader, returns error)
func (p *Product) FromJson(r io.Reader) error {
	// instantiate new encoder with the received io reader
	encoder := json.NewDecoder(r)
	// write to our product struct the decoded json (may return an error)
	return encoder.Decode(p)
}

// function that returns the product list
func GetProducts() Products {
	return productList
}

// receives a products and adds it 
func AddProduct(p *Product) {
	// assign ID of product (p) to next available int
	p.ID = getNextId()
	// append new item to productList
	productList = append(productList, p)
}

// update product with given id, may return error
func UpdateProduct(id int, p *Product) error {
	// get product, position and error from findProduct
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}
	// set product id to id received
	p.ID = id
	// update product on the list
	productList[position] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

// funct to find product in list, can return a product, position and an error
func findProduct(id int) (*Product, int, error){
	// iterate throught the product list
	for i, p := range productList {
		// if product is in the list, return it
		if p.ID == id {
			return p, i, nil
		}
	}
	// otherwise return error
	return nil, -1, ErrProductNotFound
}

// function that return a next usable id
func getNextId() int {
	// get last product of the array
	lastProduct := productList[len(productList) - 1]
	// return ID + 1 of last product
	return lastProduct.ID + 1
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