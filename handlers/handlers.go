package handlers

import "net/http"

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "index.html", nil)
}

func (a *App) About(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "about.html", nil)
}
