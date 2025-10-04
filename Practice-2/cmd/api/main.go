package main

import (
	"log"
	"mymodule/Practice-2/internal/handlers"
	"mymodule/Practice-2/internal/middleware"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/user", middleware.Auth(http.HandlerFunc(handlers.UserHandler)))

	log.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
