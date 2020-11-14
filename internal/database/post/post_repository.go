package post

import (
	"database/sql"
	"time"

	"github.com/ademalidurmus/golang-rest-api-example/internal/model"
)

// Repository ...
type Repository struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Search ...
func (r *Repository) Search() ([]model.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM posts WHERE status != 'deleted' ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var p model.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	err = rows.Err()

	return posts, err
}

// Create ...
func (r *Repository) Create(post model.Post) (model.Post, error) {
	err := r.db.QueryRow(`INSERT INTO posts (title, content, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`, post.Title, post.Content, post.Status, post.CreatedAt, post.UpdatedAt).Scan(&post.ID)
	return post, err
}

// Read ...
func (r *Repository) Read(id int) (model.Post, error) {
	var post model.Post
	err := r.db.QueryRow(`SELECT * FROM posts WHERE id=$1 AND status != 'deleted'`, id).Scan(&post.ID, &post.Title, &post.Content, &post.Status, &post.CreatedAt, &post.UpdatedAt)
	return post, err
}

// Update ...
func (r *Repository) Update(post model.Post) (model.Post, error) {
	_, err := r.db.Exec(`UPDATE posts SET title = $1, content = $2, status = $3, updated_at = $4 WHERE id = $5`, post.Title, post.Content, post.Status, post.UpdatedAt, post.ID)
	return post, err
}

// Delete ...
func (r *Repository) Delete(id int) (bool, error) {
	_, err := r.db.Exec(`UPDATE posts SET status = 'deleted', updated_at = $1 WHERE id = $2 AND status != 'deleted'`, time.Now().Format(time.RFC3339), id)
	return true, err
}
