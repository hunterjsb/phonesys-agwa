// handlers/home.go
package handlers

import (
	"agwa/models"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch some data from the model (for demonstration purposes)
	users := models.GetUsers()

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Users []models.User
	}{
		Title: "My Go Webapp",
		Users: users,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
