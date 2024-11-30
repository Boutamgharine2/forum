package main

import (
	"fmt"
	"log"
	"net/http"

	forum "forum/fonctions"

	_  "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", forum.HomeHandler)
	http.HandleFunc("/register", forum.RegisterHandl)
	http.HandleFunc("/resultat", forum.Resulfunc)
	http.HandleFunc("/login", forum.Loginhandler)
	http.HandleFunc("/resultalogin",forum.ResultaLogin)

	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("style"))))
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
