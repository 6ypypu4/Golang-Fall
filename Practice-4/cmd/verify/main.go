package main

import (
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalln("Ошибка подключения:", err)
	}
	defer db.Close()

	addUser(db)
	getAllUsers(db)
	getUserByID(db)
	updateEmail(db)
	transferBalance(db)
	deleteUser(db)
}
