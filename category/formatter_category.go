package category

import (
	"time"
)

// import "time"

type CategoryFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	CreatedAt         time.Time `json:"created_at"`
}

func FormatterCategory(category Categorys) CategoryFormatter {
	formatterCategory := CategoryFormatter{
		ID:                category.ID,
		Type:              category.Type,
		CreatedAt:         category.CreatedAt,
	}
	return formatterCategory

}

type UpdatedCategoryFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FormatterCategoryUpdated(category Categorys) UpdatedCategoryFormatter {
	formatterCategory := UpdatedCategoryFormatter{
		ID:                category.ID,
		Type:              category.Type,
		UpdatedAt:         category.UpdatedAt,
	}
	return formatterCategory

}


