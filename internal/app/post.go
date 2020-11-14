package app

import (
	"time"

	"github.com/ademalidurmus/golang-rest-api-example/internal/database/post"
	"github.com/ademalidurmus/golang-rest-api-example/internal/model"
)

// Post ...
type Post struct {
	PostRepository *post.Repository
}

// NewPostAPP ...
func NewPostAPP(p *post.Repository) Post {
	return Post{PostRepository: p}
}

// SearchPost ...
func (p *Post) SearchPost() ([]model.Post, error) {
	return p.PostRepository.Search()
}

// CreatePost ...
func (p *Post) CreatePost(post *model.Post) (model.Post, error) {
	post.CreatedAt = time.Now().Format(time.RFC3339)
	post.UpdatedAt = time.Now().Format(time.RFC3339)

	return p.PostRepository.Create(*post)
}

// ReadPost ...
func (p *Post) ReadPost(id int) (model.Post, error) {
	return p.PostRepository.Read(id)
}

// UpdatePost ...
func (p *Post) UpdatePost(post *model.Post) (model.Post, error) {
	post.UpdatedAt = time.Now().Format(time.RFC3339)

	return p.PostRepository.Update(*post)
}

// DeletePost ...
func (p *Post) DeletePost(id int) (bool, error) {
	return p.PostRepository.Delete(id)
}
