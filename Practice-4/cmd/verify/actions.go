package main

import (
	"fmt"
	"log"
	"time"

	repo "practice4-sqlx/internal/repository"

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

func addUser(db *sqlx.DB) {
	newUser := repo.User{Name: "kukich", Email: "52brat@mail.com", Balance: 75.00}
	if err := repo.InsertUser(db, newUser); err != nil {
		log.Println("Ошибка InsertUser:", err)
	}
	fmt.Println("\n\n")
}

func getAllUsers(db *sqlx.DB) {
	users, err := repo.GetAllUsers(db)
	if err != nil {
		log.Println("Ошибка GetAllUsers:", err)
	} else {
		fmt.Println("\nВсе пользователи:")
		for _, u := range users {
			fmt.Printf("\n %v %v %v %v ", u.ID, u.Name, u.Email, u.Balance)
		}
	}
	fmt.Println("\n\n")
}

func getUserByID(db *sqlx.DB) {
	user, err := repo.GetUserByID(db, 1)
	if err != nil {
		log.Println("Ошибка GetUserByID:", err)
	} else {
		fmt.Println("Пользователь с ID=1:", user)
	}
	fmt.Println("\n\n")
}

func updateEmail(db *sqlx.DB) {
	if err := repo.UpdateUserEmail(db, 2, "newbob@mail.com"); err != nil {
		log.Println("Ошибка UpdateUserEmail:", err)
	} else {
		fmt.Println(" Email обновлён успешно!")
	}
	fmt.Println("\n\n")
}

func transferBalance(db *sqlx.DB) {
	if err := repo.TransferBalance(db, 1, 2, 25.0); err != nil {
		log.Println("Ошибка TransferBalance:", err)
	} else {
		fmt.Println("Перевод успешен!")
	}
	fmt.Println("\n\n")
}

func deleteUser(db *sqlx.DB) {
	if err := repo.DeleteUser(db, 3); err != nil {
		log.Println("Ошибка DeleteUser:", err)
	} else {
		fmt.Println("Пользователь с ID=3 удалён!")
	}
}
