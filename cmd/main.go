package main

import (
	"context"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/jwt"
	"go/adv-demo/pkg/middleware"
	"net/http"
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Canceled")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}

func main2() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthRepository(userRepository)
	jwt := jwt.NewJWT(conf.Auth.Secret)

	// Handlers
	{
		auth.NewAuthHandler(router, auth.AuthHandlerDeps{
			Config:      conf,
			AuthService: authService,
			JWT:         jwt,
		})
		link.NewLinkHandler(router, link.LinkHandlerDeps{
			LinkRepository: linkRepository,
		})
	}

	// Middlwares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Server on 8081")
	server.ListenAndServe()
}
