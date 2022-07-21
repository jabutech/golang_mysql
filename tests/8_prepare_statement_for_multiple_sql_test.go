package tests

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestPrepareStatementMultipleSqlInsertData(t *testing.T) {
	// Open connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create context
	ctx := context.Background()
	// Script insert data
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	// Use Prepare statement for prepare sql
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Close prepare statement after all process end
	defer statement.Close()

	// Do Multiple insert data
	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("rizky%s@gmail.com", strconv.Itoa(i))
		comment := fmt.Sprintf("Komentar ke %s", strconv.Itoa(i))

		// Insert data with statement ExecContext (If query use QueryContext)
		result, err := statement.ExecContext(ctx, email, comment)
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
}
