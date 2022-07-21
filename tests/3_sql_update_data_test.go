package tests

import (
	"context"
	"fmt"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestSqlUpdateData(t *testing.T) {
	// Open connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create context
	ctx := context.Background()

	// Script for update
	script := "UPDATE customer SET name='joko tingkir' WHERE id='joko'"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// If success, print message to console
	fmt.Println("Update success.")
}
