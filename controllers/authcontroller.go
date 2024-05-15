package controllers

import (
	"encoding/json"
	"fmt"
	"go-marketplace/configs"
	"go-marketplace/constants"
	"go-marketplace/helpers"
	"go-marketplace/models"
	"go-marketplace/requests"
	"net/http"
	"strings"
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

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest requests.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		helpers.ResponseHelper(w, 500, err.Error(), nil)
		return
	}

	var user = models.User{}
	if err := configs.DB.First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		helpers.ResponseHelper(w, 400, "Wrong Password or Email", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, loginRequest.Password); err != nil {
		helpers.ResponseHelper(w, 404, "Wrong Password or Email", nil)
		return
	}

	token, err := helpers.CreateToken(user)
	if err != nil {
		helpers.ResponseHelper(w, 500, err.Error(), nil)
		return
	}

	helpers.ResponseHelper(w, 200, "Login Success", token)

}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get(constants.HeaderAuth)
	if tokenString == "" {
		helpers.ResponseHelper(w, http.StatusUnauthorized, "Missing authorization header", nil)
		return
	}

	if !strings.HasPrefix(tokenString, constants.TokenPrefix) {
		helpers.ResponseHelper(w, http.StatusUnauthorized, "Invalid token prefix", nil)
		return
	}
	tokenString = strings.TrimPrefix(tokenString, constants.TokenPrefix)
	err := helpers.VerifyToken(tokenString)
	if err != nil {
		helpers.ResponseHelper(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}
	fmt.Println(w, "Welcome to the the protected area")
}
