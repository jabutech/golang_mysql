package tests

import (
	"context"
	"fmt"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestSqlQueryDataWithParameter(t *testing.T) {
	// Open Connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create Context
	ctx := context.Background()

	// Sample data
	username := "admin"
	password := "admin"

	// Script query data
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	// Query data
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	// Close rows after all process scan data end
	defer rows.Close()

	// If true
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Sukses Login %s.\n", username)
	} else {
		fmt.Println("Gagal Login.")
	}
}
