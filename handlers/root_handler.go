package handlers

import "net/http"

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":   "Home Page",
		"Content": "Welcome to the home page!",
	}
	renderTemplate(w, "index.html", data)
}
