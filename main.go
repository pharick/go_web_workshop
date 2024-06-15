package main

import (
	"log"
	"net/http"

	"github.com/pharick/cool_app/handlers"
)

func main() {
	app := handlers.NewApp()

	// create a router
	r := http.NewServeMux()

	// register the handler
	r.HandleFunc("/{$}", app.Index)
	r.HandleFunc("/about", app.About)

	// start the server
	log.Printf("Server started on port 3000\n")
	http.ListenAndServe(":3000", r)
}
