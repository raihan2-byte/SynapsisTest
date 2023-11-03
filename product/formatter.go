package product

import (
	"time"
)

// import "time"

type ProductFormatter struct {
	ID int `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

func FormatterProduct(product Products) ProductFormatter {
	formatterProduct := ProductFormatter{
		ID: product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
	}
	return formatterProduct

}

type ProductGetFormatter struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatterGet(product Products) ProductGetFormatter {
	formatterGet := ProductGetFormatter{}
	formatterGet.ID = product.ID
	formatterGet.Title = product.Title
	formatterGet.Price = product.Price
	formatterGet.Stock = product.Stock
	formatterGet.CategoryID = product.CategoryID
	formatterGet.CreatedAt = product.CreatedAt

	return formatterGet
}

func FormatterGetProduct(products []Products) []ProductGetFormatter {
	productGetFormatter := []ProductGetFormatter{}

	for _, product := range products {
		productFormatter := FormatterGet(product)
		productGetFormatter = append(productGetFormatter, productFormatter)
	}

	return productGetFormatter
}

type ProductUpdateFormatter struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatterUpdate(product Products) ProductUpdateFormatter {
	formatterUpdate := ProductUpdateFormatter{}
	formatterUpdate.ID = product.ID
	formatterUpdate.Title = product.Title
	formatterUpdate.Price = product.Price
	formatterUpdate.Stock = product.Stock
	formatterUpdate.CategoryID = product.CategoryID
	formatterUpdate.CreatedAt = product.CreatedAt
	formatterUpdate.UpdatedAt = product.UpdatedAt

	return formatterUpdate
}
