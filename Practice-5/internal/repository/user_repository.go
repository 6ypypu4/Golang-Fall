package repository

import (
	"fmt"
	"strings"

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

type ProductWCategory struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	CategoryID   int    `db:"category_id"`
	Price        int    `db:"price"`
	CategoryName string `db:"category_name"`
}

// ProductFilters contains optional filters for product queries
type ProductFilters struct {
	Category string
	MinPrice *int
	MaxPrice *int
	Sort     string // "price_asc" or "price_desc"
	Limit    int
	Offset   int
}

// GetProductsWithFilters returns products with filters, sorting, and pagination
func GetProductsWithFilters(db *sqlx.DB, filters ProductFilters) ([]ProductWCategory, error) {
	var products []ProductWCategory
	var args []interface{}
	argCounter := 1

	// Build WHERE clause
	var conditions []string
	if filters.Category != "" {
		conditions = append(conditions, fmt.Sprintf("c.name = $%d", argCounter))
		args = append(args, filters.Category)
		argCounter++
	}
	if filters.MinPrice != nil {
		conditions = append(conditions, fmt.Sprintf("p.price >= $%d", argCounter))
		args = append(args, *filters.MinPrice)
		argCounter++
	}
	if filters.MaxPrice != nil {
		conditions = append(conditions, fmt.Sprintf("p.price <= $%d", argCounter))
		args = append(args, *filters.MaxPrice)
		argCounter++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Build ORDER BY clause
	orderBy := "ORDER BY p.id"
	if filters.Sort == "price_asc" {
		orderBy = "ORDER BY p.price ASC"
	} else if filters.Sort == "price_desc" {
		orderBy = "ORDER BY p.price DESC"
	}

	// Build LIMIT and OFFSET (validated to be positive integers)
	limitOffset := ""
	if filters.Limit > 0 {
		limitOffset += fmt.Sprintf(" LIMIT %d", filters.Limit)
	}
	if filters.Offset >= 0 {
		limitOffset += fmt.Sprintf(" OFFSET %d", filters.Offset)
	}

	query := fmt.Sprintf(`
		SELECT 
			p.id, 
			p.name, 
			p.category_id, 
			p.price, 
			c.name AS category_name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		%s
		%s
		%s
	`, whereClause, orderBy, limitOffset)

	err := db.Select(&products, query, args...)
	return products, err
}
