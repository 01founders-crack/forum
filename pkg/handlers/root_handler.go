package handlers

import (
	"forum/pkg/auth"
	"forum/pkg/db"
	"forum/pkg/models"
	"log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
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

	likeCounts := make(map[int]int)

	// Loop through all posts and count the number of likes for each post
	for _, post := range allPosts {
		postID := post.ID
		likeCount, err := db.CountLikesByPost(postID)
		if err != nil {
			// Handle the error if necessary
			log.Println("Error counting likes for post", postID, ":", err)
			continue // Continue to the next iteration
		}
		likeCounts[postID] = likeCount
	}

	// User is authenticated
	// Pass the session data to the template
	data := map[string]interface{}{
		"Title":       "Home Page",
		"SessionData": session.Values,
		"IsHomePage":  true,
		"Posts":       allPosts,
		"Categories":  allCategories,
		"Likes":       likeCounts,
	}
	renderTemplate(w, "index.html", data)
}
