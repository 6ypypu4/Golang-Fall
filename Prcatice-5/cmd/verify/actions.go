package main

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "user=postgres password=postgres dbname=usersdb sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}

/*
func addUser(db *sqlx.DB) {
	newUser := repo.User{Name: "kukich", Email: "52brat@mail.com", Balance: 75.00}
	if err := repo.InsertUser(db, newUser); err != nil {
		log.Println("Ошибка InsertUser:", err)
	}
	fmt.Println("\n\n")
}

*/
