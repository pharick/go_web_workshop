package main

import (
	"log"

	"github.com/pharick/cool_app/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	// register the handler
	a.RegisterHandler("/{$}", a.Index)
	a.RegisterHandler("/about", a.About)
	a.RegisterHandler("/users", a.Users)

	// start the server
	a.Serve(3000)
}
