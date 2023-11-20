package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var MyDBVar *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	// Open the SQLite database file
	MyDBVar, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Could not open database: ", err)
	}

	// Create Users Table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err = MyDBVar.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Could not create users table: ", err)
	}

	// Create Posts Table
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		img VARCHAR(255),
		category_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (category_id) REFERENCES categories (id)
	);`
	_, err = MyDBVar.Exec(createPostsTable)
	if err != nil {
		log.Fatal("Could not create posts table: ", err)
	}

	// Create Comments Table
	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		content TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`
	_, err = MyDBVar.Exec(createCommentsTable)
	if err != nil {
		log.Fatal("Could not create comments table: ", err)
	}

	// Create Categories Table
	createCategoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	);`
	_, err = MyDBVar.Exec(createCategoriesTable)
	if err != nil {
		log.Fatal("Could not create categories table: ", err)
	}

	// Create Likes Table
	createLikesTable := `
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`
	_, err = MyDBVar.Exec(createLikesTable)
	if err != nil {
		log.Fatal("Could not create likes table: ", err)
	}

	// Create DisLikes Table
	createDisLikesTable := `
	CREATE TABLE IF NOT EXISTS dislikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`
	_, err = MyDBVar.Exec(createDisLikesTable)
	if err != nil {
		log.Fatal("Could not create dislikes table: ", err)
	}
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return MyDBVar
}
