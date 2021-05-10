package controllers

import (
	"log"
	"net/http"
	"store/models"
	"strconv"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	templ.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error on price convertion")
		}

		quantityConvToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error on quantity convertion")
		}

		models.CreateProduct(name, description, priceConvToFloat, quantityConvToInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
