package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"jabutech.com/golang_mysql/config"
)

func TestSqlQueryData(t *testing.T) {
	// Open Connection
	db := config.GetConnection()
	// Close connection after all process done
	defer db.Close()

	// Create Context
	ctx := context.Background()

	// Script query data
	script := `SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer`
	// Query data
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Close rows after all process scan data end
	defer rows.Close()

	// Do iteration all rows if data is available (true)
	for rows.Next() {
		var id, name string
		var email sql.NullString // In MySQL type data string but is Null
		var balance int32
		var rating float64        // In MySQL type data DOUBLE
		var birtDate sql.NullTime // In MySQL type data DATE but is Null
		var createdAt time.Time
		var married bool

		// Scan every rows, and passing into variable in top
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birtDate, &createdAt, &married)
		if err != nil {
			panic(err)
		}

		// Print data to console
		fmt.Printf("id: %s\n", id)
		fmt.Printf("name: %s\n", name)
		// Check is email Valid (true) or is available, print email
		if email.Valid {
			fmt.Printf("email: %v\n", email.String)
		}
		fmt.Printf("balance: %v\n", balance)
		fmt.Printf("rating: %.1f\n", rating)
		// Check is birthDate Valid (true) or is available, print birthDate
		if birtDate.Valid {
			fmt.Printf("rating: %v\n", birtDate.Time)
		}
		fmt.Printf("created at: %s\n", createdAt)
		fmt.Printf("married: %t\n", married)
		fmt.Println("=========")
	}
}
