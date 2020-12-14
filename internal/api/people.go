package api

import (
	"encoding/json"
	"net/http"

	"github.com/ademalidurmus/golang-rest-api-example/internal/app"
	"github.com/ademalidurmus/golang-rest-api-example/internal/model"
)

// PeopleAPI ...
type PeopleAPI struct {
	Person app.Person
}

// NewPeopleAPI ...
func NewPeopleAPI(p app.Person) PeopleAPI {
	return PeopleAPI{Person: p}
}

// Encrypt ...
func (p PeopleAPI) Encrypt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person model.Person

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			RespondWithError(w, http.StatusNotAcceptable, err.Error())
			return
		}

		person = p.Person.Encrypt(person)

		RespondWithJSON(w, http.StatusCreated, person)
	}
}
