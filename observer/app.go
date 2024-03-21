package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"company.com/observer/products"
)

type Application struct{}

type MyObserver struct{}

func (o *MyObserver) ProductAdded(product products.Product) {
	log.Info().Msg(fmt.Sprintf("Product added %v", product))
}

type MyOtherObserver struct{}

func (o *MyOtherObserver) ProductAdded(product products.Product) {
	log.Info().Msg(fmt.Sprintf("Product added %v from MyOtherObserver", product))
}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	var productList products.ProductList = products.NewProductList()
	var myObserver products.ProductListObserver = &MyObserver{}
	otherObserver := &MyOtherObserver{}

	productList.AddObserver(myObserver)
	productList.AddObserver(otherObserver)

	var product = products.NewProduct("Apple")

	productList.AddProduct(product)
}
