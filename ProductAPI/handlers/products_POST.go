package handlers

import (
	data "github.com/ramanchawla1290/GO/ProductAPI/data"

	"net/http"
)

//	swagger:route POST /products products addProduct
//
//	Creates a new product in the database
//
//		Schemes: http
//
//		Responses:
//		201: productResponse
//		400: errorInvalidJSON
//		422: errorValidation
//		500: errorServer
//
//		Produces:
//  	- application/json
//
//		Consumes:
//		- application/json
//
//		Deprecated: false

func (ph *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	// AddProduct creates a new product in the database
	prod := r.Context().Value("prod").(*data.Product)

	err := data.AddProduct(prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		prod.ID = -1
		return
	}
}
