package handlers

import (
	"forum/pkg/db"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HandleRegisterAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordSpace := false

	for _, char := range password {
		if unicode.IsSpace(int32(char)) {
			passwordSpace = true
		}
	}

	if db.CheckUserExists(email, username) {
		http.Error(w, "User Already Exists", http.StatusInternalServerError)
		return
	}

	if passwordSpace {
		http.Error(w, "Password cannot contain spaces!", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusInternalServerError)
		return
	}

	_, err = db.AddUser(email, username, string(hashedPassword))
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
