package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalln("Ошибка подключения:", err)
	}
	defer db.Close()

	GetAllProducts(db)
	//FilterByCategory(db, "Electric")
}
