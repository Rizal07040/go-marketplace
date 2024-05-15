package routes

import (
	"github.com/gorilla/mux"
	"go-marketplace/controllers/v1"
)

func UserRouter(controller *controllers.UserController, r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.HandleFunc("/list-all", controller.GetAllUser).Methods("GET")

}
