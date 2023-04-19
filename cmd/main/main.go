package main

import (
	"github_takuma411/Golang-MySQL-CRUD-Bookstore-Management-API-master/pkg/routes"
	"log"
	"net/http"
)
import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
