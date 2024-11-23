package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	createDatabase()
}

func createDatabase() {
	// Ouvrir la base de données SQLite
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Lire le fichier SQL pour créer les tables
	schema, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Exécuter le script SQL pour créer les tables
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de données et tables créées avec succès.")
}

func Registre(r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmplt/register.html"))
	r.ParseForm()
	email := r.Form.Get("email")
	if !chackemail(email) {
		



	}
}
