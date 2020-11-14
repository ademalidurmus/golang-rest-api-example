package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/ademalidurmus/golang-rest-api-example/internal/api"
	"github.com/ademalidurmus/golang-rest-api-example/internal/database"
)

func main() {
	conn, err := database.DBConn(os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_PORT"), os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASS"), os.Getenv("APP_DB_NAME"))

	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: api.NewRouter(conn),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
