package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	router    *http.ServeMux
	templates *Templates
	// will add database connection
	// will add any other dependencies
}

func NewApp() *App {
	return &App{
		router:    http.NewServeMux(),
		templates: NewTemplates(),
	}
}

func (a *App) RegisterHandler(url string, handler http.HandlerFunc) {
	a.router.HandleFunc(url, handler)
}

func (a *App) Serve(port int) {
	log.Printf("Server started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.router)
}
