package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID      int
	Text    string
	Created time.Time
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=db user=root password=root dbname=next_go sslmode=disable")
	checkErr(err)
}

func main() {
	defer db.Close()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var f Todo
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&f)
		checkErr(err)
	}

	switch r.URL.Path {
	case "/create":
		create(f.Text)
	case "/read":
		fmt.Println("Read")
	case "/update":
		update(f.ID, f.Text)
	case "/delete":
		delete(f.ID)
	default:
		fmt.Println("Noting")
	}
	todo := read()
	js, err := json.Marshal(todo)
	checkErr(err)
	w.Write(js)
}

func create(text string) {
	stmt, err := db.Prepare("INSERT INTO todos (text,created) VALUES ($1,$2) RETURNING id")
	checkErr(err)
	_, err = stmt.Exec(text, time.Now())
	checkErr(err)
}

func read() []Todo {
	rows, err := db.Query("SELECT * FROM todos")
	checkErr(err)
	var todo []Todo
	for rows.Next() {
		var id int
		var text string
		var created time.Time
		err = rows.Scan(&id, &text, &created)
		checkErr(err)
		todo = append(todo, Todo{id, text, created})
	}
	return todo
}

func update(id int, text string) {
	stmt, err := db.Prepare("UPDATE todos SET text=$1 WHERE id=$2")
	checkErr(err)
	_, err = stmt.Exec(text, id)
	checkErr(err)
}

func delete(id int) {
	stmt, err := db.Prepare("DELETE FROM todos WHERE id=$1")
	checkErr(err)
	_, err = stmt.Exec(id)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
