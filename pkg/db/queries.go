package db

import (
	"database/sql"
	"fmt"
	"forum/pkg/models"
)

// AddUser adds a new user to the users table
func AddUser(email, username, password string) (sql.Result, error) {
	query := `INSERT INTO users (email, username, password) VALUES (?, ?, ?)`
	return MyDBVar.Exec(query, email, username, password)
}

// GetUserByEmail retrieves user details by email
func GetUserByEmail(email string) (*sql.Row, error) {
	query := `SELECT * FROM users WHERE email = ?`
	return MyDBVar.QueryRow(query, email), nil
}

func CheckUserExists(email, username string) bool {
	query := "SELECT COUNT(*) FROM users WHERE email = ? OR username = ?"

	var count int
	err := MyDBVar.QueryRow(query, email, username).Scan(&count)
	if err != nil {
		fmt.Println("Error checking user existence:", err)
		return false
	}

	return count > 0
}

// AddPost adds a new post to the posts table
func AddPost(userID, categoryID int, content string, title string, img string) (sql.Result, error) {
	query := `INSERT INTO posts (user_id, category_id, content, title, img) VALUES (?, ?, ?, ?, ?)`
	return MyDBVar.Exec(query, userID, categoryID, content, title, img)
}

// GetPostsByCategory retrieves posts by category
func GetPostsByCategory(categoryID int) ([]models.Post, error) {
	query := `SELECT * FROM posts WHERE category_id = ?`
	rows, err := MyDBVar.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Img, &post.CategoryID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetAllPosts retrieves all posts
func GetAllPosts() ([]models.Post, error) {
	query := `SELECT * FROM posts`
	rows, err := MyDBVar.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Img, &post.CategoryID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// AddComment adds a new comment to the comments table
func AddComment(userID, postID int, content string) (sql.Result, error) {
	query := `INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)`
	return MyDBVar.Exec(query, userID, postID, content)
}

// GetCommentsByPost retrieves comments by post
func GetCommentsByPost(postID int) (*sql.Rows, error) {
	query := `SELECT * FROM comments WHERE post_id = ?`
	return MyDBVar.Query(query, postID)
}

// AddLike adds a new like to the likes table
func AddLike(userID, postID int) (sql.Result, error) {
	query := `INSERT INTO likes (user_id, post_id) VALUES (?, ?)`
	return MyDBVar.Exec(query, userID, postID)
}

// CountLikesByPost counts the number of likes for a specific post
func CountLikesByPost(postID int) (*sql.Row, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = ?`
	return MyDBVar.QueryRow(query, postID), nil
}

// AddCategory adds a new category to the categories table
func AddCategory(name string) (sql.Result, error) {
	query := `INSERT INTO categories (name) VALUES (?)`
	return MyDBVar.Exec(query, name)
}

// GetAllCategories retrieves all categories
func GetAllCategories() ([]models.Category, error) {
	rows, err := MyDBVar.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name) // replace with your actual scan
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryById retrieves a category by id
func GetCategoryIDByName(name string) (int, error) {
	var id int
	err := MyDBVar.QueryRow("SELECT id FROM categories WHERE name = ?", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
