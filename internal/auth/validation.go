package auth

import (
	"encoding/json"
	"go/adv-demo/pkg"
	"net/http"
)

func LoginBodyValidation(body LoginRequest, w http.ResponseWriter) bool {
	if body.Email == "" || body.Password == "" {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(pkg.StatusCode["BAD_REQUEST"])
		json.NewEncoder(w).Encode(pkg.ValidationError{
			Code:    pkg.StatusCode["BAD_REQUEST"],
			Message: "email or password should not be empty",
		})
		return false
	}
	return true
}
