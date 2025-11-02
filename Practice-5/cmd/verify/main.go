package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"practice5-sqlx/internal/handlers"
)

func main() {
	// Connect to PostgreSQL
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=usersdb sslmode=disable host=localhost port=5432")
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer db.Close()

	log.Println("Connected to database successfully")

	// Register routes
	http.HandleFunc("/products", handlers.GetProductsHandler(db))

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
