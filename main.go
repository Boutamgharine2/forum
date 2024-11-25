package main

import (
	"fmt"
	"log"
	"net/http"

	forum "forum/fonctions"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", forum.Homhandler)
	http.HandleFunc("/register", forum.RegisterHandl)
	http.HandleFunc("/resultat",forum.Resulfunc)
	//forum.CreateDatabase()
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
