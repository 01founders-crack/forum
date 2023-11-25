package handlers

import (
	"forum/pkg/auth"
	"forum/pkg/db"
	"forum/pkg/models"
	"net/http"
)

func HandleCategory(w http.ResponseWriter, r *http.Request) {
	// You can retrieve the post data from the database here
	session, err := auth.Store.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check all errors here and handle them i didnt do it @@@@
	var allCategories []models.Category
	allCategories, _ = db.GetAllCategories()

	// Check all errors here and handle them i didnt do it @@@@
	var allPosts []models.Post
	allPosts, _ = db.GetAllPosts()

	// User is authenticated
	// Pass the session data to the template
	data := map[string]interface{}{
		"Title":       "Home Page",
		"SessionData": session.Values,
		"IsHomePage":  true,
		"Posts":       allPosts,
		"Categories":  allCategories,
	}
	renderTemplate(w, "index.html", data)
}
