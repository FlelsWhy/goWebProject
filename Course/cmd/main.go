package main

import (
	"COURSE/pkg/api"
	"COURSE/pkg/config"
	"COURSE/pkg/repositori"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	db, err := repositori.New(config.DBADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	api := api.New(mux.NewRouter(), db)
	api.Endpoints()
	log.Fatal(api.ListenAndServe(config.PORT))
}
