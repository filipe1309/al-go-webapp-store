package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectDatabase() *sql.DB {
	connection := "user=postgres dbname=go_store password=postgres host=store sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDatabase()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-shirt", Description: "Blue, very beatiful", Price: 39, Quantity: 5},
		{"Shoes", "Confortable", 89, 3},
		{"Headphode", "Very good", 59, 2},
		{"New Prodcut", "Very nice", 99, 1},
	}
	templ.ExecuteTemplate(w, "Index", products)
}
