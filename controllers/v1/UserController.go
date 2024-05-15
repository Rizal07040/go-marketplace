package controllers

import (
	"go-marketplace/controllers"
	"go-marketplace/helpers"
	"go-marketplace/service/v1"
	"net/http"
	"net/http/httptest"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userControllerImpl *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	response := httptest.NewRecorder()
	controllers.ProtectedHandler(response, r)

	if response.Code != http.StatusOK {
		// If ProtectedHandler returns an error status, copy the response and return
		for k, v := range response.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(response.Code)
		w.Write(response.Body.Bytes())
		return
	}
	data, err := userControllerImpl.userService.GetAllUsers()
	if err != nil {
		helpers.ResponseHelper(w, http.StatusBadRequest, "Users is Empty", nil)
		return
	}
	helpers.ResponseHelper(w, 200, "Users retrieved successfully", data)
}
