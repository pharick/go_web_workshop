package main

import (
	"github.com/pharick/cool_app/app"
)

func main() {
	a := app.NewApp()

	// register the handler
	a.RegisterHandler("/{$}", a.Index)
	a.RegisterHandler("/about", a.About)
	a.RegisterHandler("/users", a.Users)
	a.RegisterHandler("/users/{username}", a.User)

	// start the server
	a.Serve(3000)
}
