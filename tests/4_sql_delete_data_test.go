package tests

import (
	"context"
	"fmt"
	"testing"

	"jabutech.com/golang_mysql/config"
)

func TestSqlDeleteData(t *testing.T) {
	// Open connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create context
	ctx := context.Background()

	// Script delete
	script := `DELETE FROM customer WHERE id="joko"`
	// Update
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// If success, print message to console
	fmt.Println("Delete success.")
}
