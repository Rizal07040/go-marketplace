package controllers

import (
	"encoding/json"
	"go-marketplace/configs"
	"go-marketplace/helpers"
	"go-marketplace/models"
	"go-marketplace/requests"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest requests.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		helpers.ResponseHelper(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if registerRequest.Password != registerRequest.PasswordConfirm {
		helpers.ResponseHelper(w, 400, "Password not match", nil)
		return
	}

	passwordHash, err := helpers.HashPassword(registerRequest.Password)
	if err != nil {
		helpers.ResponseHelper(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.ResponseHelper(w, 500, err.Error(), nil)
	}

	helpers.ResponseHelper(w, 200, "Register Success", nil)

}
