package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare(`INSERT INTO usuarios (nome) VALUES (?)`)
	stmt.Exec("Mari")
	stmt.Exec("João")

	result, _ := stmt.Exec("Pedro")
	id, _ := result.LastInsertId()
	fmt.Println(id)

	linhas, _ := result.RowsAffected()
	fmt.Println(linhas)
}
