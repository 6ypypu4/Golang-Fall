package repository

import (
	"github.com/jmoiron/sqlx"
)

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Product struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	CategoryID int    `db:"category_id"`
	Price      int    `db:"price"`
}

func GetProducts(db *sqlx.DB) ([]Product, error) {
	var products []Product
	err := db.Select(&products, "SELECT * FROM products ORDER BY id")
	return products, err
}
