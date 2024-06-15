package main

import (
	"github.com/pharick/cool_app/handlers"
)

func main() {
	app := handlers.NewApp()

	// register the handler
	app.RegisterHandler("/{$}", app.Index)
	app.RegisterHandler("/about", app.About)
	app.RegisterHandler("/users", app.Users)
	app.RegisterHandler("/users/{username}", app.User)

	// start the server
	app.Serve(3000)
}
