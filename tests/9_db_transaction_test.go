package tests

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestTransaction(t *testing.T) {
	// Open Connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create context
	ctx := context.Background()
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// Script insert data
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("rizky%s@gmail.com", strconv.Itoa(i))
		comment := fmt.Sprintf("Komentar ke %s", strconv.Itoa(i))

		// Insert data with statement ExecContext (If query use QueryContext)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		// Get last insert id
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		// If success, print log
		fmt.Println("Comment Id", id)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	// If process fail can use `tx.Rollback()` as cancellation all sql
}
