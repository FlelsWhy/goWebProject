package api

import (
	"log"
	"net/http"
)

func (api *api) Midllware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}
