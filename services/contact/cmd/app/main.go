package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })
	// fmt.Println("Server listening on port 8080")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("Server error:", err)
	// }

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "godb"
	)
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database")
}
