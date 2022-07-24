package main

import (
	"context"
	"log"
	"main/src/microservices/restful-services/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
	// "./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create a "new products" handler
	ph := handlers.NewProducts(l)

	// create new http serve mux
	sm := http.NewServeMux()
	// assign hh handler to "/" path on new surve mux
	sm.Handle("/", ph)
	
	// create a new instance of http Server
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	// start a goroutine 
	go func() {
		// listen and serve our server
		err := s.ListenAndServe() 
		if(err != nil){
			l.Fatal(err)
		}
	}()

	// shut down server gracefully when program gets killed
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	// pass the reference of sigCahn to sig
	sig := <- sigChan
	l.Println("Graceful shutdown", sig)

	// create an instance of a shutdown with a deadline of 30s
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// wait until there is no work being done and shutdown gracefully
	s.Shutdown(tc)
}