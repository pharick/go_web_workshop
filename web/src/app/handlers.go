package app

import "net/http"

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "index.html", map[string]any{
		"title":  "Title from data",
		"header": "Header from data",
	})
}

func (a *App) About(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "about.html", nil)
}

type User struct {
	Id       int
	Username string
	Email    string
}

// create a slice (array) of users
var users = []User{
	{Id: 1, Username: "john", Email: "john@example.com"},   // 0
	{Id: 2, Username: "jane", Email: "jane@example.com"},   // 1
	{Id: 3, Username: "mary", Email: "mary@example.com"},   // 2
	{Id: 4, Username: "peter", Email: "peter@example.com"}, // 3
}

func (a *App) Users(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "users.html", map[string]any{
		"users": users,
	})
}

func (a *App) User(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	index := -1
	for i, user := range users {
		if user.Username == username {
			index = i
			break
		}
	}

	// if the user is not found
	if index < 0 {
		http.NotFound(w, r)
		return
	}

	a.templates.Render(w, "user.html", map[string]any{
		"user": users[index],
	})
}
