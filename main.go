package main

import (
	"github.com/gorilla/mux"
	"go-marketplace/configs"
	"go-marketplace/routes"
	"log"
	"net/http"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api/v1").Subrouter()

	routes.AuthRouter(router)

	log.Println("Server running on Port 8080")
	http.ListenAndServe(":8080", router)
}
