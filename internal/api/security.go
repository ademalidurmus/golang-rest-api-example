package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var replacer = strings.NewReplacer("\\u003e", "", "\\u003c", "", ">", "", "<", "")

// XSSSecurityMiddleware ...
func XSSSecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data interface{}

		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			_ = json.Unmarshal(body, &data)

			m, _ := data.(map[string]interface{})
			iteratedData := iterate(m)
			jsonBytes, err := json.Marshal(iteratedData)
			if err == nil {
				r.Body = ioutil.NopCloser(strings.NewReader(replacer.Replace(string(jsonBytes))))
			}
		}

		next.ServeHTTP(w, r)
	})
}

// https://stackoverflow.com/questions/48949737/how-to-use-reflect-to-recursively-parse-nested-struct-in-go
func iterate(data interface{}) interface{} {
	d := reflect.ValueOf(data)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		returnSlice := make([]interface{}, d.Len())
		for i := 0; i < d.Len(); i++ {
			returnSlice[i] = iterate(d.Index(i).Interface())
		}
		return returnSlice
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		tmpData := make(map[string]interface{})
		for _, k := range d.MapKeys() {
			tmpData[k.String()] = iterate(d.MapIndex(k).Interface())
		}
		return tmpData
	} else {
		return data
	}
}
