package auth

import (
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
	*jwt.JWT
}
type AuthHandler struct {
	*configs.Config
	*AuthService
	*jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
		JWT:         deps.JWT,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := handler.JWT.Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := LoginResponse{
			Token: token,
		}

		res.Json(w, response, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := handler.JWT.Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := LoginResponse{
			Token: token,
		}

		res.Json(w, response, http.StatusOK)
	}
}
