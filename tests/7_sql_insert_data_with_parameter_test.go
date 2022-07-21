package tests

import (
	"context"
	"fmt"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestSqlInsertDataWithParameter(t *testing.T) {
	// Open connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Sample data
	username := "letenk"
	password := "admin"

	// Create context
	ctx := context.Background()

	// Script insert data to table user
	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	// If success, print log
	fmt.Println("Success create user.")
}
