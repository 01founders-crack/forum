package handlers

import (
	"fmt"
	"forum/pkg/db"
	"forum/pkg/models"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	// You can retrieve the post data from the database here
	postID, err := extractID(r.URL.Path)
	if err != nil {
		// Handle the error, for example, return a 400 Bad Request
		http.Error(w, "Invalid ID in URL", http.StatusBadRequest)
		return
	}
	fmt.Println("Post ID:", postID)
	var post models.Post
	post, err = db.GetPostByID(postID)
	if err != nil {
		// Handle the error, for example, return a 404 Not Found
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Title":   "Sample Post Title",
		"Content": "This is a sample post content.",
		"Post":    post,
	}
	renderTemplate(w, "post.html", data)
}
