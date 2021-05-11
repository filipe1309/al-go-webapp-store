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

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	templ.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error on id convertion")
		}

		priceConvToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error on price convertion")
		}

		quantityConvToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error on quantity convertion")
		}

		models.UpdateProduct(idConvToInt, name, description, priceConvToFloat, quantityConvToInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
