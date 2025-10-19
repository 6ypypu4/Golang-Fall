package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Email   string  `db:"email"`
	Balance float64 `db:"balance"`
}

func InsertUser(db *sqlx.DB, user User) error {
	query := `INSERT INTO users (name, email, balance) VALUES (:name, :email, :balance)`
	_, err := db.NamedExec(query, user)
	return err
}

func GetAllUsers(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users ORDER BY id")
	return users, err
}

func GetUserByID(db *sqlx.DB, id int) (User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return user, err
}

func UpdateUserEmail(db *sqlx.DB, id int, newEmail string) error {
	query := `UPDATE users SET email=$1 WHERE id=$2`
	_, err := db.Exec(query, newEmail, id)
	return err
}

func DeleteUser(db *sqlx.DB, id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}

func TransferBalance(db *sqlx.DB, fromID int, toID int, amount float64) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	var senderBalance float64
	err = tx.Get(&senderBalance, "SELECT balance FROM users WHERE id=$1", fromID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("отправитель не найден")
	}
	if senderBalance < amount {
		tx.Rollback()
		return fmt.Errorf("недостаточно средств")
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id=$2", amount, fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id=$2", amount, toID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("получатель не найден")
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
