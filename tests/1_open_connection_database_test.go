package tests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"jabutech.com/golang_mysql/config"
)

func TestOpenConnectionToMySQL(t *testing.T) {
	// Open connection use function db
	db := config.GetConnection()

	// Close db after use
	defer db.Close()
}
