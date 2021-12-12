package handlers

import (
	"net/http"
	"strconv"

	data "github.com/ramanchawla1290/GO/ProductAPI/data"

	mux "github.com/gorilla/mux"
)

//	swagger:route PUT /products/{id} products updateProduct
//
//	Updates the product with the specified Id in the database with new product data
//
//		Schemes: http
//
//		Responses:
//		201: productResponse
//		400: errorInvalidJSON
//		404: errorProductNotFound
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

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// UpdateProduct modifies the product at the specified id
	prod := r.Context().Value("prod").(*data.Product)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ph.logger.Println("ERROR converting Id:", params["id"])
		http.Error(w, "ERROR processing Product Id", http.StatusUnprocessableEntity)
		prod.ID = -1
		return
	}

	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		prod.ID = -1
		return
	}
}
