package main

import (
	"fmt"
	"github.com/fernandes-dev/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Printf("🚀 Server started at http://localhost:3333\n")
	log.Fatal(http.ListenAndServe("localhost:3333", r))
}
