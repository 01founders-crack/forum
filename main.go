package main

import (
	"fmt"
	"forum/pkg/db"
	"forum/pkg/handlers"
	"net/http"
)

func main() {
	db.InitDB()
	fmt.Println("db created")
	// Create a new ServeMux (request multiplexer)
	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handler for the home page
	mux.HandleFunc("/", handlers.HandleRoot)

	// Handler for the login page
	mux.HandleFunc("/login", handlers.HandleLogin)

	// Handler for the registration page
	mux.HandleFunc("/register", handlers.HandleRegister)

	// Handler for the registration authentication
	mux.HandleFunc("/registerauth", handlers.HandleRegisterAuth)

	// Handler for the registration page
	mux.HandleFunc("/create_post", handlers.HandleCreatePost)

	// Handler for a sample post page
	mux.HandleFunc("/post/1", handlers.HandlePost)

	// Create an HTTP server with the chosen address and ServeMux
	server := &http.Server{
		Addr:    ":8080", // Listen on port 8080
		Handler: mux,     // Use the ServeMux as the handler
	}

	// Start the HTTP server
	fmt.Println("Server is running on : http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
