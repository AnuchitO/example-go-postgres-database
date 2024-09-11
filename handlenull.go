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

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status INTEGER 
	);
	`
	_, err = db.Exec(createTb) // HL

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", "iPad Pro", nil) // HL
	var id int
	err = row.Scan(&id) // HL
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)

	sol1(db)
	sol3(db)
	sol2(db)
}

// SOL 1  OMIT
func sol1(db *sql.DB) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1") // HL
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title string
	var status *string // SOLUTION : #1 // HL

	err = row.Scan(&id, &title, &status) // HL
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one row", id, title, status)
}

// ESOL 1 OMIT

// SOL 3 OMIT
func sol3(db *sql.DB) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1") // HL
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId) // HL
	var id int
	var title string
	var status sql.NullString // SOLUTION : #2 // HL

	err = row.Scan(&id, &title, &status) // HL
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one row", id, title, status.String)
}

// ESOL 3 OMIT

// SOL 2 OMIT
type MaybeString string

func (n *MaybeString) Scan(value interface{}) error {
	if value == nil {
		*n = ""
		return nil
	}

	switch s := value.(type) {
	case string:
		*n = MaybeString(s)
		return nil
	default:
		return fmt.Errorf("can't convert: %T to type int32", value)
	}
}

func sol2(db *sql.DB) {
	row := stmt.QueryRow(rowId) // HL
	var id int
	var title string
	var status MaybeString // SOLUTION : #2 // HL

	err = row.Scan(&id, &title, &status) // HL

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1") // HL
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1

	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one row", id, title, status)
}

// ESOL 2 OMIT
