package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare(`UPDATE usuarios SET nome = ? WHERE id = ?`)
	stmt.Exec("Maria", 1)
	stmt.Exec("Joana", 2)
	stmt.Exec("Pedro", 3)

	stmt2, _ := db.Prepare(`DELETE FROM usuarios WHERE id = ?`)
	stmt2.Exec(3)
}
