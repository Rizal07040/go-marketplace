package routes

import (
	"github.com/gorilla/mux"
	"go-marketplace/controllers"
)

func AuthRouter(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	//router.HandleFunc("/login", controllers.Login).Methods("POST")
}
