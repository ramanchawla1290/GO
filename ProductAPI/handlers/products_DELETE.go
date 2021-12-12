package handlers

import (
	"net/http"
	"strconv"

	data "github.com/ramanchawla1290/GO/ProductAPI/data"

	mux "github.com/gorilla/mux"
)

//	swagger:route DELETE /products/{id} products deleteProduct
//
//  Deletes a product with the specified id from the database
//
//		Schemes: http
//
//		Responses:
//		204: successNoContent
//		404: errorProductNotFound
//		422: errorValidation
//		500: errorServer
//
//		Produces:
//  	- application/json
//
//		Deprecated: false

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// DeleteProduct removes the product with the specified id from the database
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ph.logger.Println("ERROR converting Id:", params["id"])
		http.Error(w, "Unable to convert ID", http.StatusUnprocessableEntity)
		return
	}

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
