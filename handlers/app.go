package handlers

import "net/http"

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
