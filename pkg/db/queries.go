package db

import (
	"database/sql"
)

// AddUser adds a new user to the users table
func AddUser(email, username, password string) (sql.Result, error) {
	query := `INSERT INTO users (email, username, password) VALUES (?, ?, ?)`
	return db.Exec(query, email, username, password)
}

// GetUserByEmail retrieves user details by email
func GetUserByEmail(email string) (*sql.Row, error) {
	query := `SELECT * FROM users WHERE email = ?`
	return db.QueryRow(query, email), nil
}

// AddPost adds a new post to the posts table
func AddPost(userID, categoryID int, content string) (sql.Result, error) {
	query := `INSERT INTO posts (user_id, category_id, content) VALUES (?, ?, ?)`
	return db.Exec(query, userID, categoryID, content)
}

// GetPostsByCategory retrieves posts by category
func GetPostsByCategory(categoryID int) (*sql.Rows, error) {
	query := `SELECT * FROM posts WHERE category_id = ?`
	return db.Query(query, categoryID)
}

// AddComment adds a new comment to the comments table
func AddComment(userID, postID int, content string) (sql.Result, error) {
	query := `INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)`
	return db.Exec(query, userID, postID, content)
}

// GetCommentsByPost retrieves comments by post
func GetCommentsByPost(postID int) (*sql.Rows, error) {
	query := `SELECT * FROM comments WHERE post_id = ?`
	return db.Query(query, postID)
}

// AddLike adds a new like to the likes table
func AddLike(userID, postID int) (sql.Result, error) {
	query := `INSERT INTO likes (user_id, post_id) VALUES (?, ?)`
	return db.Exec(query, userID, postID)
}

// CountLikesByPost counts the number of likes for a specific post
func CountLikesByPost(postID int) (*sql.Row, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = ?`
	return db.QueryRow(query, postID), nil
}

// AddCategory adds a new category to the categories table
func AddCategory(name string) (sql.Result, error) {
	query := `INSERT INTO categories (name) VALUES (?)`
	return db.Exec(query, name)
}

// GetAllCategories retrieves all categories
func GetAllCategories() (*sql.Rows, error) {
	query := `SELECT * FROM categories`
	return db.Query(query)
}
