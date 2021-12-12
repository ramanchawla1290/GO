package data

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	validator "github.com/go-playground/validator"
)

// 	Structure of the API Product
// 	swagger:model
type Product struct {
	// 	Product ID
	ID int `json:"id"`

	// 	Name of the Product
	// 		min length: 3
	//		max length: 20
	Name string `json:"name" validate:"required"`

	// 	Description of the Product
	// 		min length: 10
	//		max length: 50
	Description string `json:"description" validate:"required"`

	// 	Price of the Product
	// 		min: 0.01
	Price float64 `json:"price" validate:"gte=0.01"`

	// 	SKU of the Product
	//		min length: 8
	//		max length: 15
	SKU string `json:"sku" validate:"required,sku"` // Custom Validaor 'sku'

	//	Creation Time of the Product
	CreateOn string `json:"-"` //omitting field in json

	// Last Update Time of the Product
	UpdateOn string `json:"-"` //omitting field in json
}

// Custom Validator Function for SKU
func validateSKU(fl validator.FieldLevel) bool {
	// Considering SKU format as "abc-abcd-abcdef"
	re := regexp.MustCompile(`[A-Z]+-[A-Z]+-[0-9]+`)
	result := re.FindAllString(fl.Field().String(), -1)

	return len(result) == 1 // true only if one sku
}

// Validator Function - USING GO VALIDATOR
func (prod *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(prod) // validationErrors = err.(validator.ValidationErrors)
}

// Method to read JSON from an io.Reader and Convert it into 'Product' object
func (prod *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(prod)
}

// Method to convert 'Product' object into JSON and write on an io.Writer
func (prod *Product) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(prod)
}

// Custom type to store slice of 'Product' type objects
type ProductList []*Product

// Method to convert 'ProductList' object into JSON and write on an io.Writer
func (pl *ProductList) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(pl)
}

var prodList ProductList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "COF-WHT-001",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "COF-BLK-001",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}

func getNextId() int {
	if len(prodList) == 0 {
		return 1
	}
	lastProd := prodList[len(prodList)-1]
	return lastProd.ID + 1
}

func findProduct(id int) (index int) {
	for index, prod := range prodList {
		if prod.ID == id {
			return index
		}
	}
	return -1
}

func GetProducts() *ProductList {
	return &prodList
}

func GetProduct(id int) (*Product, error) {
	i := findProduct(id)
	if i < 0 {
		return nil, &ProductError{"Product ID Not Found"}
	}

	return prodList[i], nil
}

func AddProduct(prod *Product) error {
	prod.ID = getNextId()
	prod.CreateOn = time.Now().UTC().String()
	prod.UpdateOn = time.Now().UTC().String()

	prodList = append(prodList, prod)
	return nil
}

func UpdateProduct(id int, prod *Product) error {
	i := findProduct(id)
	if i < 0 {
		return &ProductError{"Product ID Not Found"}
	}

	prod.ID = id
	prod.CreateOn = prodList[i].CreateOn
	prod.UpdateOn = time.Now().UTC().String()
	prodList[i] = prod

	return nil
}

func DeleteProduct(id int) error {
	i := findProduct(id)
	if i < 0 {
		return &ProductError{"Product ID Not Found"}
	}

	prodList = append(prodList[:i], prodList[i+1:]...)
	return nil
}

type ProductError struct {
	errorMessage string
}

func (err *ProductError) Error() string {
	return err.errorMessage
}
