package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ademalidurmus/golang-rest-api-example/internal/app"
	"github.com/ademalidurmus/golang-rest-api-example/internal/model"
	"github.com/gorilla/mux"
)

// PostAPI ...
type PostAPI struct {
	Post app.Post
}

// NewPostAPI ...
func NewPostAPI(p app.Post) PostAPI {
	return PostAPI{Post: p}
}

// SearchPost ...
func (p PostAPI) SearchPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := p.Post.SearchPost()
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, posts)
	}
}

// CreatePost ...
func (p PostAPI) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post model.Post

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&post)
		if err != nil {
			RespondWithError(w, http.StatusNotAcceptable, err.Error())
			return
		}

		post, err = p.Post.CreatePost(&post)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusCreated, post)
	}
}

// ReadPost ...
func (p PostAPI) ReadPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		post, err := p.Post.ReadPost(id)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, post)
	}
}

// UpdatePost ...
func (p PostAPI) UpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		post, err := p.Post.ReadPost(id)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&post)
		if err != nil {
			RespondWithError(w, http.StatusNotAcceptable, err.Error())
			return
		}

		post.ID = id

		post, err = p.Post.UpdatePost(&post)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, post)
	}
}

// DeletePost ...
func (p PostAPI) DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		post, err := p.Post.ReadPost(id)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		_, err = p.Post.DeletePost(id)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusNoContent, post)
	}
}
