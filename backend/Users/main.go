package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "bubble.db.elephantsql.com"
	port     = 5432
	user     = "wzpwshep"
	password = "Lk6rq_iCBbYLk-LxtTI8uIpM7Myee4Rt"
	dbname   = "wzpwshep"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := sql.Open("postgres", connStr)

	return &Database{
		db: db,
	}
}

func (d *Database) Close() {
	d.db.Close()
}

func (db *Database) Registration(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	login := r.FormValue("login")
	password := r.FormValue("password")

	query := fmt.Sprintf("INSERT INTO users (name, login, password) VALUES ('%s', '%s', '%s')", name, login, password)
	_, err := db.db.Exec(query)

	if err != nil {
		fmt.Fprintf(w, "%s", "Error executing query")
		return
	}

	fmt.Fprintf(w, "%s", "Success")
}

func (db *Database) Authorisation(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	var result string
	query := fmt.Sprintf("SELECT name FROM users WHERE login='%s' AND password='%s'", login, password)
	err := db.db.QueryRow(query).Scan(&result)
	if err != nil {
		fmt.Fprintf(w, "%s", "Invalid login or password")
		return

	} else {
		fmt.Fprintf(w, "%s", "Valid login and password")
	}

}

func (db *Database) Sale(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")

	var result string
	query := fmt.Sprintf("SELECT total_sum FROM users WHERE login='%s'", login)
	err := db.db.QueryRow(query).Scan(&result)
	if err != nil {
		fmt.Fprintf(w, "%s", "Error executing query")
		return
	}
	fmt.Println(result)
	res, _ := strconv.ParseFloat(result, 10)
	var sale string
	if res <= 100000 {
		sale = fmt.Sprintf("%.2f", res*0.000001)
	} else {
		sale = "0.1"
	}
	fmt.Fprintf(w, "%s", sale)
}

func (db *Database) Bought(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	sum, _ := strconv.Atoi(r.FormValue("sum"))

	query := fmt.Sprintf("UPDATE users SET total_sum = total_sum + %d WHERE login='%s'", sum, login)
	res, err := db.db.Exec(query)
	rowsAffected, err1 := res.RowsAffected()
	if err != nil || rowsAffected == 0 || err1 != nil {
		fmt.Fprintf(w, "%s", "Error")
		return
	}

	fmt.Fprintf(w, "%s", "Success")
}

func (db *Database) Delete(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")

	query := fmt.Sprintf("DELETE from users WHERE login='%s'", login)
	res, err := db.db.Exec(query)
	rowsAffected, err1 := res.RowsAffected()
	if err != nil || rowsAffected == 0 || err1 != nil {
		fmt.Fprintf(w, "%s", "Error")
		return
	}

	fmt.Fprintf(w, "%s", "Success")
}

func (db *Database) ChangePw(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	query := fmt.Sprintf("UPDATE users SET password='%s'WHERE login='%s'", password, login)
	res, err := db.db.Exec(query)
	rowsAffected, err1 := res.RowsAffected()
	if err != nil || rowsAffected == 0 || err1 != nil {
		fmt.Fprintf(w, "%s", "Error")
		return
	}

	fmt.Fprintf(w, "%s", "Success")
}

func (db *Database) ChangeName(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	name := r.FormValue("name")

	query := fmt.Sprintf("UPDATE users SET name='%s'WHERE login='%s'", name, login)
	res, err := db.db.Exec(query)
	rowsAffected, err1 := res.RowsAffected()
	if err != nil || rowsAffected == 0 || err1 != nil {
		fmt.Fprintf(w, "%s", "Error")
		return
	}

	fmt.Fprintf(w, "%s", "Success")
}

func (db *Database) Name(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")

	var result string
	query := fmt.Sprintf("SELECT name FROM users WHERE login='%s'", login)
	err := db.db.QueryRow(query).Scan(&result)
	if err != nil {
		fmt.Fprintf(w, "%s", "Error executing query")
		return
	}
	fmt.Println(result)
	fmt.Fprintf(w, "%s", result)
}

func (db *Database) CheckStatus(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")

	var result string
	query := fmt.Sprintf("SELECT status FROM users WHERE login='%s'", login)
	err := db.db.QueryRow(query).Scan(&result)
	if err != nil {
		fmt.Fprintf(w, "%s", "Error executing query")
		return
	}
	fmt.Println(result)
	fmt.Fprintf(w, "%s", result)
}

func SetupServer(db *Database) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/registration", db.Registration)
	mux.HandleFunc("/authorisation", db.Authorisation)
	mux.HandleFunc("/sale", db.Sale)
	mux.HandleFunc("/bought", db.Bought)
	mux.HandleFunc("/delete", db.Delete)
	mux.HandleFunc("/change_pw", db.ChangePw)
	mux.HandleFunc("/change_name", db.ChangeName)
	mux.HandleFunc("/name", db.Name)
	mux.HandleFunc("/check_status", db.CheckStatus)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func main() {
	db := NewDatabase()
	defer db.Close()

	server := SetupServer(db)

	log.Fatal(server.ListenAndServe())
}