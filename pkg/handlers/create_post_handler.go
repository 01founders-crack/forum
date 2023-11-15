package handlers

import (
	"forum/pkg/auth"
	"net/http"
)

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	// You can retrieve the post data from the database here
	session, err := auth.Store.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// User is authenticated
	// Pass the session data to the template
	data := map[string]interface{}{
		"Title":       "Sample create post Title",
		"SessionData": session.Values,
	}

	renderProfileTemplate(w, "create_post.html", data)
}
