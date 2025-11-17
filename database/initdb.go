package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	// Initialize database connection here
	host := "localhost"
	port := "5432"
	user := "example"
	pass := "example"
	dbname := "mydb"
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, sslmode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Println("Database connection established")
	return db, nil
}