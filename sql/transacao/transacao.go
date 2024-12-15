package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(`INSERT INTO usuarios (id, nome) VALUES (?, ?)`)
	stmt.Exec(2000, "Mari")
	stmt.Exec(2001, "João")
	_, err = stmt.Exec(2001, "Pedro")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
