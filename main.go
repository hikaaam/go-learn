package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllerHome "github.com/hikaaam/go-learn/src/controller/home"
	"github.com/hikaaam/go-learn/src/model"
)

func main() {
	// run database migrations
	conn := model.CreateConnection()
	model.Migrations(conn)

	//API router
	r := mux.NewRouter()
	r.HandleFunc("/", controllerHome.GetGorillas).Methods("GET")
	r.HandleFunc("/{id}", controllerHome.ShowGorilla).Methods("GET")
	r.HandleFunc("/add", controllerHome.AddGorrila).Methods("POST")
	r.HandleFunc("/edit", controllerHome.EditGorilla).Methods("PUT")
	r.HandleFunc("/delete/{id}", controllerHome.DeleteGorilla).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
