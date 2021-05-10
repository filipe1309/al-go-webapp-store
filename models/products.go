package models

import "store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.ConnectDatabase()

	selectAllProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insertDbData, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertDbData.Exec(name, description, price, quantity)

	defer db.Close()
}
