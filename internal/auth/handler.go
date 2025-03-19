package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg"
	"net/http"
	"regexp"
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
		match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, payload.Email)
		if !match {
			pkg.Json(w, "Email is wrong", pkg.StatusCode["BAD_REQUEST"])
			return
		}
		if payload.Password == "" {
			pkg.Json(w, "Password is required", pkg.StatusCode["BAD_REQUEST"])
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
