package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/ademalidurmus/golang-rest-api-example/internal/database"
)

type db struct {
	conn *sql.DB
}

type post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	conn, err := database.DBConn(os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_PORT"), os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASS"), os.Getenv("APP_DB_NAME"))

	if err != nil {
		panic(err)
	}

	db := &db{conn: conn}

	router := mux.NewRouter()
	router.HandleFunc("/posts", db.getPosts).Methods("GET")
	router.HandleFunc("/posts", db.createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", db.getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", db.updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", db.deletePost).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func (db *db) getPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.conn.Query(`SELECT * FROM posts WHERE status != 'deleted' ORDER BY id DESC`)
	if err != nil {
		response(w, error{Message: err.Error(), Status: "error", Code: http.StatusInternalServerError})
		return
	}
	defer rows.Close()

	var posts []post

	for rows.Next() {
		var p post

		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			response(w, error{Message: err.Error(), Status: "error", Code: http.StatusInternalServerError})
			return
		}

		posts = append(posts, p)
	}

	err = rows.Err()
	switch {
	case err != nil:
		response(w, error{Message: err.Error(), Status: "error", Code: http.StatusInternalServerError})
	default:
		if posts == nil {
			response(w, make([]int, 0))
			return
		}
		response(w, posts)
	}
}

func (db *db) createPost(w http.ResponseWriter, r *http.Request) {
	var p post

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		response(w, error{Message: "Invalid request body", Status: "error", Code: http.StatusNotAcceptable})
		return
	}

	p.CreatedAt = time.Now().Format(time.RFC3339)
	p.UpdatedAt = time.Now().Format(time.RFC3339)

	err = db.conn.QueryRow(`INSERT INTO posts (title, content, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`, p.Title, p.Content, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&p.ID)
	switch {
	case err != nil:
		response(w, error{Message: err.Error(), Status: "error", Code: http.StatusInternalServerError})
	default:
		response(w, p)
	}
}

func (db *db) getPost(w http.ResponseWriter, r *http.Request) {
	var p post

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response(w, error{Message: "Unexpected post id", Status: "error", Code: http.StatusNotAcceptable})
		return
	}

	err = db.conn.QueryRow(`SELECT * FROM posts WHERE id=$1 AND status != 'deleted'`, id).Scan(&p.ID, &p.Title, &p.Content, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		response(w, error{Message: "Post not found", Status: "error", Code: http.StatusNotFound})
	default:
		response(w, p)
	}
}

func (db *db) updatePost(w http.ResponseWriter, r *http.Request) {
	var p post

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response(w, error{Message: "Unexpected post id", Status: "error", Code: http.StatusNotAcceptable})
		return
	}

	err = db.conn.QueryRow(`SELECT * FROM posts WHERE id = $1 AND status != 'deleted'`, id).Scan(&p.ID, &p.Title, &p.Content, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		response(w, error{Message: "Post not found", Status: "error", Code: http.StatusNotFound})
	default:
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&p)
		if err != nil {
			response(w, nil)
			return
		}

		p.UpdatedAt = time.Now().Format(time.RFC3339)

		_, err = db.conn.Exec(`UPDATE posts SET title = $1, content = $2, status = $3, updated_at = $4 WHERE id = $5`, p.Title, p.Content, p.Status, p.UpdatedAt, p.ID)
		switch {
		case err != nil:
			response(w, error{Message: err.Error(), Status: "error", Code: http.StatusInternalServerError})
		default:
			response(w, p)
		}
	}
}

func (db *db) deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response(w, error{Message: "Unexpected post id", Status: "error", Code: http.StatusNotAcceptable})
		return
	}

	res, err := db.conn.Exec(`UPDATE posts SET status = 'deleted', updated_at = $1 WHERE id = $2 AND status != 'deleted'`, time.Now().Format(time.RFC3339), id)
	switch {
	case err != nil:
		response(w, error{Message: err.Error(), Status: "error", Code: http.StatusInternalServerError})
	default:
		count, countErr := res.RowsAffected()
		if countErr != nil {
			response(w, error{Message: countErr.Error(), Status: "error", Code: http.StatusInternalServerError})
			return
		} else if count == 0 {
			response(w, error{Message: "Post not found", Status: "error", Code: http.StatusNotFound})
		}
		response(w, nil)
	}
}

func response(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}
