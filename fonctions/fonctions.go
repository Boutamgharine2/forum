package forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func Homhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// ErrorHandler(w, r, http.StatusNotFound, "page Not found")
		//	fmt.Println("hiii")
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not alowed", http.StatusMethodNotAllowed)
	}
	tmpl, err := (template.ParseFiles("template/index.html"))
	r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrrrrrrrra!")
		return
	}
}

func Resulfunc(w http.ResponseWriter, r *http.Request) {
	// type Message struct {
	// 	message string
	// }
	var Message string

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := (template.ParseFiles("template/result.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	r.ParseForm()
	email := r.Form.Get("email")
	username := r.Form.Get("username")
	password := r.Form.Get("mypassword")
	confpassword := r.Form.Get("confirmation")

	if password != confpassword {
		Message="les mots de passe sont pas identiques !"
		tmpl.Execute(w, Message)
		return

	}

	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		fmt.Println("hiiii")
		log.Fatal(err)
	}
	defer db.Close()
	v, err := (emailExiste(db, email))
	if v && err == nil {
		Message="l'email que tu utuluser est deja exist!"
		tmpl.Execute(w, Message)

		return
	}

	if email != "" && username != "" && password != "" {
		err0 := inseredata(db, email, username, password)
		if err0 != nil {
			Message="l'operation est echoe"
			tmpl.Execute(w, Message)

			return
		}
		tmpl.Execute(w, "toust est bien!!")
		return

	}
	tmpl.Execute(w, "quelque chose n'est pas correct!")
}

func RegisterHandl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := (template.ParseFiles("template/register.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	r.ParseForm()

	// 	email := r.Form.Get("email")
	// 	username := r.Form.Get("username")
	// 	password := r.Form.Get("password")
	// 	confpassword := r.Form.Get("confirmation")

	// 	if password != confpassword {
	// 		fmt.Println("les mots de passe sont pas corespondant!")
	// 		return
	// 	}

	// 	db, err := sql.Open("sqlite3", "./mydatabase.db")
	// 	if err != nil {
	// 		fmt.Println("hiiii")
	// 		log.Fatal(err)
	// 	}
	// 	defer db.Close()
	// 	v, err := (emailExiste(db, email))
	// 	if v && err == nil {
	// 		http.Error(w, "deja exist", http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	if email!= "" && username != "" &&password != "" {
	// 	err0 := inseredata(db, email, username, password)
	// 	if err0 != nil {
	// 		fmt.Println(err0)
	// 		return
	// 	}
	// }

	// supprimerUtilisateur(db, 1)
	tmpl.Execute(w, 0)
}

func supprimerUtilisateur(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := db.Exec(query, id)
	return err
}

func inseredata(db *sql.DB, email, username, password string) error {
	query := "INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)"

	_, err := db.Exec(query, email, username, password)
	return err
}

func emailExiste(db *sql.DB, email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ?"

	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func CreateDatabase() {
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de données et tables créées avec succès.")
}
