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

func (a *App) Users(w http.ResponseWriter, r *http.Request) {
	res, err := a.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []User
	// Iterate over the rows of sql query result
	// and append each user to the users slice
	for res.Next() {
		var user User
		res.Scan(&user.Id, &user.Username, &user.Email)
		users = append(users, user)
	}

	a.templates.Render(w, "users.html", map[string]any{
		"users": users,
	})
}
