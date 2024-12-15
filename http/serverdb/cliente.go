package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarios(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpe :(")
	}
}

func usuarioPorID(w http.ResponseWriter, _ *http.Request, id int) {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(json))
}

func usuarios(w http.ResponseWriter, _ *http.Request) {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, u)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(json))
}
