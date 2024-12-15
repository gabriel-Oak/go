package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id < ?", 6)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u.id, u.nome)
	}
}
