package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// START OMIT
func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", "buy bmw", "active") // HL
	var id int
	err = row.Scan(&id) // HL
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}

	fmt.Println("insert todo success id : ", id)
}

// END OMIT
