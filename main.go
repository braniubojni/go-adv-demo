package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server on 8081")
	server.ListenAndServe()
}
