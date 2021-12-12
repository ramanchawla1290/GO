package handlers

import (
	"net/http"
	"strconv"

	data "github.com/ramanchawla1290/GO/ProductAPI/data"

	mux "github.com/gorilla/mux"
)

// Ping function responds to ping requests from clients at the API BasePath "/"
func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// 	swagger:route GET /products products listProducts
//
// 	Returns a list of all products in the database
//
//  	Schemes: http
//
//     	Responses:
//      200: productList
//		500: errorServer
//
//     	Produces:
//     	- application/json
//
//     	Deprecated: false

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// GetAllProducts returns all the products from the data store
	prodList := data.GetProducts()

	w.Header().Set("Content-Type", "application/json")

	err := prodList.ToJSON(w)
	if err != nil {
		ph.logger.Println("ERROR serializing product json")
		http.Error(w, "ERROR sending json response", http.StatusInternalServerError)
		return
	}
}

// 	swagger:route GET /products/{id} products getProduct
//
// 	Returns a single product from the database
//
//  	Schemes: http
//
//     	Responses:
//      200: productResponse
//		404: errorProductNotFound
//		422: errorValidation
//		500: errorServer
//
//     	Produces:
//     	- application/json
//
//     	Deprecated: false

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// GetProduct returns a single product from the data store
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ph.logger.Println("ERROR converting Id:", params["id"])
		http.Error(w, "Unable to convert ID", http.StatusUnprocessableEntity)
		return
	}

	prod, err := data.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = prod.ToJSON(w)
	if err != nil {
		ph.logger.Println("ERROR serializing product json")
		http.Error(w, "ERROR sending json response", http.StatusInternalServerError)
		return
	}
}
