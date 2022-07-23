package main

import (
	"log"
	"net/http"
	"os"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create a "new hello" handler
	hh := handlers.NewHello(l)
	// create a "new goodbye" handler
	gh := handlers.NewGoodbye(l)

	// create new http serve mux
	sm := http.NewServeMux()
	// assign hh handler to "/" path on new surve mux
	sm.Handle("/", hh)
	// assign gh handler to "/goodbye" path on new surve mux
	sm.Handle("/goodbye", gh)
	
	// serve on port 9090 specifying our serve mux
	http.ListenAndServe(":9090", sm)
}