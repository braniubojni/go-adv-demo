package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	// Handlers
	{
		auth.NewAuthHandler(router, auth.AuthHandlerDeps{
			Config: conf,
		})
		link.NewLinkHandler(router, link.LinkHandlerDeps{
			Config: conf,
		})
	}

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server on 8081")
	server.ListenAndServe()
}
