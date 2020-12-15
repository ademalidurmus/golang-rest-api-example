package api

import (
	"database/sql"
	"net/http"

	"github.com/ademalidurmus/xsscleaner"
	"github.com/gorilla/mux"

	"github.com/ademalidurmus/golang-rest-api-example/internal/app"
	"github.com/ademalidurmus/golang-rest-api-example/internal/database/post"
)

// Router ...
type Router struct {
	router *mux.Router
	db     *sql.DB
}

// NewRouter ...
func NewRouter(db *sql.DB) *Router {
	r := &Router{
		router: mux.NewRouter(),
		db:     db,
	}
	r.initRoutes()
	return r
}

// ServeHTTP ...
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) initRoutes() {
	postAPI := InitPostAPI(r.db)
	r.router.HandleFunc("/posts", postAPI.SearchPost()).Methods(http.MethodGet)
	r.router.HandleFunc("/posts", postAPI.CreatePost()).Methods(http.MethodPost)
	r.router.HandleFunc("/posts/{id:[0-9]+}", postAPI.ReadPost()).Methods(http.MethodGet)
	r.router.HandleFunc("/posts/{id:[0-9]+}", postAPI.UpdatePost()).Methods(http.MethodPut)
	r.router.HandleFunc("/posts/{id:[0-9]+}", postAPI.DeletePost()).Methods(http.MethodDelete)

	peopleAPI := InitPeopleAPI()
	r.router.HandleFunc("/people/_encrypt", peopleAPI.Encrypt()).Methods(http.MethodPost)

	r.router.Use(xsscleaner.Middleware)
}

// InitPostAPI ..
func InitPostAPI(db *sql.DB) PostAPI {
	postRepository := post.NewRepository(db)
	postAPP := app.NewPostAPP(postRepository)
	return NewPostAPI(postAPP)
}

// InitPeopleAPI ..
func InitPeopleAPI() PeopleAPI {
	peopleAPP := app.NewPeopleAPP()
	return NewPeopleAPI(peopleAPP)
}
