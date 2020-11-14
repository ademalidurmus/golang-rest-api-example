package database

import "github.com/ademalidurmus/golang-rest-api-example/internal/model"

// PostRepository interface is the common interface for a repository
type PostRepository interface {
	Search() ([]model.Post, error)
	Create(model.Post) (model.Post, error)
	Read(id int) (model.Post, error)
	Update(model.Post) (model.Post, error)
	Delete(id int) (bool, error)
}
