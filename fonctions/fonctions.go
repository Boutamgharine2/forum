package forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := (template.ParseFiles("template/Login.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	r.ParseForm()
	tmpl.Execute(w, 0)
}

func Resulfunc(w http.ResponseWriter, r *http.Request) {
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
		Message = "the passwords are not the same!"
		tmpl.Execute(w, Message)
		return

	}

	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		Message = "Internal server error!"
		tmpl.Execute(w, Message)
		return

	}
	defer db.Close()
	v, err := (EmailOrUsernameExiste(db, email, username))
	if v && err == nil {
		Message = "The email or username you are using already exists!"
		tmpl.Execute(w, Message)

		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if email != "" && username != "" && string(hashedPassword) != "" {
		err0 := inseredata(db, email, username, string(hashedPassword))
		if err0 != nil {
			Message = "the operation failed!"
			tmpl.Execute(w, Message)

			return
		}
		tmpl.Execute(w, "your operation is successful")
		return

	}
	tmpl.Execute(w, "something is not right!")
}

func ResultaLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "methode notAlowed", http.StatusMethodNotAllowed)
		return

	}
	tmpl, err := template.ParseFiles("./template/resultalogin.html")
	if err != nil {
		http.Error(w, "Internal server Error!", http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	email := r.Form.Get("email0")

	password := r.Form.Get("mypassword0")
	fmt.Println(email, password)
	tmpl.Execute(w, "hi")
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

func EmailOrUsernameExiste(db *sql.DB, email, username string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ? OR username = ?"

	var count int
	err := db.QueryRow(query, email, username).Scan(&count)
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

	fmt.Println("Database and tables created successfully.")
}
