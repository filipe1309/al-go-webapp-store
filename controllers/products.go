package controllers

import (
	"net/http"
	"store/models"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	templ.ExecuteTemplate(w, "Index", allProducts)
}
