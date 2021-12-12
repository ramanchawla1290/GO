// 		Product API
//
//		The purpose of this applciation is to provide a fully functional prototype of
//		a REST API developed using GO and related frameworks
//
//		This API is also swagger 2.0 compliant.
//
// 		Schemes: http
//		Host: localhost:9090
// 		BasePath: /
// 		Version: 0.0.1
//
// 		Consumes:
// 		- application/json
//
// 		Produces:
// 		- application/json
//
// swagger:meta

package handlers

import (
	data "github.com/ramanchawla1290/GO/ProductAPI/data"
)

// 	Documentation

//swagger:model
type productRequest struct {

	// 	Name of the product
	// 		required: true
	// 		min length: 3
	//		max length: 20
	Name string `json:"name"`

	// 	Description of the product
	//		required:true
	//		min length: 10
	//		max length: 50
	Description string `json:"description"`

	//	Price of the product
	//		required:true
	//		min : 0.01
	Price float64 `json:"price"`

	//	SKU of the product
	//		required:true
	//		pattern: [A-Z]+-[A-Z]+-[0-9]+
	//		min length: 8
	//      max length: 15
	SKU string `json:"sku"`
}

// 	swagger:parameters addProduct updateProduct
type productRequestWrapper struct {

	//	Product Information required to create / update a product
	//		in: body
	Body productRequest `json:"body"`
}

// 	swagger:parameters getProduct updateProduct deleteProduct
type productIdWrapper struct {
	// 	The Product ID to update / delete
	// 		required: true
	// 		min: 1
	//		in: path
	ID int `json:"id"`
}

// 	A list of products
// 	swagger:response productList
type productListWrapper struct {
	// in: body
	Body []data.Product `json:"body"`
}

// 	Details of the product created/updated in the database
// 	swagger:response productResponse
type productWrapper struct {
	// in: body
	Body data.Product `json:"body"`
}

//  No Content
// 	swagger:response successNoContent
type noContentWrapper struct {
}

// 	Error message as a string response for invalid JSON payload
// 	swagger:response errorInvalidJSON
type errorBadRequestWrapper struct {
	//	in: body
	Body string `json:"body"`
}

// 	Error message as a string response for Product Id not found
// 	swagger:response errorProductNotFound
type errorProductNotFoundWrapper struct {
	//	in: body
	Body string `json:"body"`
}

// 	Error message as a string response for failed Validation
// 	swagger:response errorValidation
type errorValidationWrapper struct {
	//	in: body
	Body string `json:"body"`
}

// 	Error message as a string response for internal server error
// 	swagger:response errorServer
type errorServerWrapper struct {
	//	in: body
	Body string `json:"body"`
}
