package config

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// Open connection
	db, err := sql.Open("mysql", "secret:secret@tcp(localhost:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	// (3) Set koneksi pertama yang dibuat saat aplikasi dijalankan
	db.SetMaxIdleConns(10)
	// (4) Set maksimal berapa banyak koneksi yang dibatasi ketika ada request open koneksi masuk
	db.SetMaxOpenConns(100)
	// (5) Set durasi koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxIdleTime(5 * time.Minute)
	// (6) Set duration untuk memperbaharui seluruh koneksi dengan yang baru
	db.SetConnMaxLifetime(60 * time.Minute)

	// Return connection for use other function
	return db
}
