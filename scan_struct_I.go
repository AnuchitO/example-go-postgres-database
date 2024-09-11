package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	// START OMIT
	type Todo struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1") // HL
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string

	err = row.Scan(&id, &title, &status) // HL
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	t := &Todo{ID: id, Title: title, Status: status} // HL

	fmt.Printf("%#v\n", t)
	// END OMIT
}
