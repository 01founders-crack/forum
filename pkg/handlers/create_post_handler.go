package handlers

import "net/http"

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	// You can retrieve the post data from the database here
	data := map[string]interface{}{
		"Title":   "Sample create post Title",
		"Content": "This is a sample post content.",
	}
	renderTemplate(w, "create_post.html", data)
}
