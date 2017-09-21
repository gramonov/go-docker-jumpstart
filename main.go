// RESTful API layer for realtime meeting notes
// Works in conjunction with our single page app

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	var (
		port = flag.String("port", "3000", "Port to listen on")
		host = flag.String("host", "", "Address to listen on")
	)
	flag.Parse()

	// Unfortunately, at the time "realize" (live-reload tool we're
	// using) doesn't seem to properly work with command line args, so
	// for now we provide a way to override flags with environment
	// variables
	if envar := os.Getenv("APP_PORT"); envar != "" {
		*port = envar
	}
	if envar := os.Getenv("APP_HOST"); envar != "" {
		*host = envar
	}

	addr := *host + ":" + *port
	log.Printf("Starting server on address %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
