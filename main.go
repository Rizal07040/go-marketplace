package main

import (
	"github.com/gorilla/mux"
	"go-marketplace/configs"
	"go-marketplace/controllers/v1"
	"go-marketplace/models/repositories"
	"go-marketplace/routes"
	"go-marketplace/service/v1/impl"
	"log"
	"net/http"
)

func main() {
	configs.ConnectDB()

	//Init Repository
	userRepository := repositories.UserRepository{configs.DB}

	//Init Service
	userService := service.NewUserServiceImpl(userRepository)

	//Init controller
	userController := controllers.NewUserController(userService)

	r := mux.NewRouter()
	router := r.PathPrefix("/api/v1").Subrouter()

	routes.AuthRouter(router)
	routes.UserRouter(userController, router)

	log.Println("Server running on Port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
