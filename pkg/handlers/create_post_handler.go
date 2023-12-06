package handlers

import (
	"forum/pkg/auth"
	"forum/pkg/db"
	"net/http"
)

func HandleCreatePostPage(w http.ResponseWriter, r *http.Request) {
	// You can retrieve the post data from the database here
	session, err := auth.Store.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := db.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		// Handle form submission
		r.ParseForm()
		title := r.Form.Get("title")
		content := r.Form.Get("content")
		category := r.Form.Get("category")
		categoryID, _ := db.GetCategoryIDByName(category)
		img := ""
		userID, ok := session.Values["user_id"].(int)
		if !ok {
			// Handle error: user ID is not an integer or is nil
			http.Error(w, "Session expired", http.StatusUnauthorized)
			return
		}
		if content == "" || title == "" || category == "" {
			// If the form submission is invalid, set an error message
			session.Values["notification"] = "Invalid form submission"
			// Save the session
			// session.Save(r, w)

			// Redirect back to the create post page
			http.Redirect(w, r, "/create_post", http.StatusSeeOther)
			return

		}

		if r.Form.Get("img") != "" {
			img = r.Form.Get("img")
			// !!!!!!!check here and add img to db add also into ui side@@@@@@ !!!!!!!!!
		}
		result, err := db.AddPost(userID, categoryID, content, title, img)
		if err != nil || result == nil {
			// If the operation failed, set an error message
			session.Values["notification"] = "Failed to create post"
		} else {
			// If the operation was successful, set a success message
			session.Values["notification"] = "Post created successfully"
		}

		// Save the session
		session.Save(r, w)

		// Redirect back to the create post page
		http.Redirect(w, r, "/create_post", http.StatusSeeOther)
	} else {
		// Render the form
		data := map[string]interface{}{
			"Title":        "Sample create post Title",
			"SessionData":  session.Values,
			"IsCreatePost": true,
			"Categories":   categories,
		}

		renderTemplate(w, "create_post.html", data)
	}
}
