package handlers

import "net/http"

func HandlePost(w http.ResponseWriter, r *http.Request) {
	// You can retrieve the post data from the database here
	data := map[string]interface{}{
		"Title":   "Sample Post Title",
		"Content": "This is a sample post content.",
	}
	renderTemplate(w, "post.html", data)
}
