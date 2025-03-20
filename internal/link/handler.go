package link

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
)

type LinkHandlerDeps struct {
	*configs.Config
}
type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateLinkRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)

		response := &CreateLinkResponse{
			Success: true,
		}

		res.Json(w, response, pkg.StatusCode["SUCCESS"])
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
