package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	// Set the response content type
	w.Header().Set("Content-Type", "text/html")

	// Parse header, footer, and page templates
	tmplFiles := []string{
		"templates/layout.html", // Ensure "layout.html" is executed first
		"templates/components/head.html",
		"templates/components/header.html",
		"templates/components/footer.html",
		"templates/pages/" + tmpl,
	}

	// Parse the templates
	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the templates with data
	err = t.ExecuteTemplate(w, "layout", data) // Execute "layout" template first
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderLogRegTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	// Set the response content type
	w.Header().Set("Content-Type", "text/html")

	// Parse header, footer, and page templates
	tmplFiles := []string{
		"templates/layout_logreg.html", // Ensure "layout.html" is executed first
		"templates/components/head_logreg.html",
		"templates/components/header_logreg.html",
		"templates/pages/" + tmpl,
	}

	// Parse the templates
	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the templates with data
	err = t.ExecuteTemplate(w, "layout_logreg", data) // Execute "layout" template first
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Rin add layout_profile.html head_profile.html  header_profile.html
func renderProfileTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	// Set the response content type
	w.Header().Set("Content-Type", "text/html")

	// Parse header, footer, and page templates
	tmplFiles := []string{
		"templates/layout_profile_own.html", // Ensure "layout.html" is executed first
		"templates/components/head_profile.html",
		"templates/components/header_profile.html",
		"templates/pages/" + tmpl,
	}

	// Parse the templates
	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the templates with data
	err = t.ExecuteTemplate(w, "layout_profile_own", data) // Execute "layout" template first
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]

	// Convert the string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Return an error if the conversion fails
		return 0, err
	}

	return id, nil
}
