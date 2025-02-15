package api

import (
	"COURSE/pkg/repositori"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repositori.PGRepo
}

func New(r *mux.Router, db *repositori.PGRepo) *api {
	return &api{r: r, db: db}
}

func (api *api) Endpoints() {
	api.r.HandleFunc("/api/hello", api.hello).Queries("id", "{id}")
	api.r.HandleFunc("/api/hello", api.hello)
	api.r.Use(api.Midllware)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
