package tests

import (
	"context"
	"fmt"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestSqlInsertData(t *testing.T) {
	// Open connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create context
	ctx := context.Background()

	// Create script sql
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('rizky', 'Rizky Darmawan', 'letenk@gmail.com', 1000000, 5.0, '1994-3-23', true), ('dwi', 'Dwi Aprilia', 'dwi@gmail.com', 2000000, 5.0, '1994-4-12', true), ('aisyah', 'Hadzkya Aisyah', NULL, 500000, 5.0, NULL, true)"
	// Execute insert data
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// If success, print info
	fmt.Println("Success insert new customer.")
}
