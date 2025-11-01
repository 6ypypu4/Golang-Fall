package repository

import "github.com/jmoiron/sqlx"

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

type ProductWCategory struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	CategoryID   int    `db:"category_id"`
	Price        int    `db:"price"`
	CategoryName string `db:"category_name"`
}

func GetAllProductsWithCategory(db *sqlx.DB) ([]ProductWCategory, error) {
	var products []ProductWCategory
	query := `
		SELECT 
			p.id, 
			p.name, 
			p.category_id, 
			p.price, 
			c.name AS category_name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		ORDER BY p.id;
	`
	err := db.Select(&products, query)
	return products, err
}
