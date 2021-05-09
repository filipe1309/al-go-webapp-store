package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
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
