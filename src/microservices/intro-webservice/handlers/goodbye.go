package handlers

import (
	"log"
	"net/http"
)

// struct that contains log class
type Goodbye struct {
	log *log.Logger
}

// function that gets a logger and returns Goodbye handler
func NewGoodbye(log *log.Logger) *Goodbye{
	return &Goodbye{log}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.log.Println("Goodbye World")
}