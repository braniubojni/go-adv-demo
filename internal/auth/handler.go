package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			pkg.Json(w, err.Error(), pkg.StatusCode["BAD_REQUEST"])
			return
		}
		isLoggedIn := LoginBodyValidation(payload, w)
		if !isLoggedIn {
			return
		}

		res := LoginResponse{
			Token: "123",
		}

		pkg.Json(w, res, pkg.StatusCode["SUCCESS"])
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
