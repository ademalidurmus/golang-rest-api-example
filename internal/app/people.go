package app

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"

	"github.com/ademalidurmus/golang-rest-api-example/internal/model"
)

var customTagName = "adem"

// Person ...
type Person model.Person

// NewPeopleAPP ...
func NewPeopleAPP() Person {
	return Person{}
}

// Encrypt ...
func (p *Person) Encrypt(person model.Person) model.Person {
	newPerson := encryptTagsValue(person)
	return newPerson
}

func encryptTagsValue(person model.Person) model.Person {
	t := reflect.TypeOf(person)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(customTagName)

		if tag == "true" {
			v := reflect.ValueOf(&person).Elem().FieldByIndex(field.Index).String()
			reflect.ValueOf(&person).Elem().FieldByIndex(field.Index).SetString(getMD5Hash(v))
		}
	}

	return person
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	_, _ = hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
