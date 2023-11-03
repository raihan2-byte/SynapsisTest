package handler

import (
	"net/http"
	"strconv"
	"synapsisid/helper"
	"synapsisid/product"
	"synapsisid/user"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.ServiceProduct
}

func NewProductHandler(service product.ServiceProduct) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := product.FormatterProduct(newProduct)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetAllProduct (c *gin.Context){
	products, err := h.productService.GetProducts()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, product.FormatterGetProduct(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	products, err := h.productService.GetProduct(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID product.GetinputProductID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Product with that id not found")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData product.UpdatedProduct

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "invalid input data")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	newProduct, err := h.productService.UpdatedProduct(inputID, inputData)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to update product")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, product.FormatterUpdate(newProduct))
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var input product.GetinputProductID

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)
	// userID := currentUser.ID

	newDel, err := h.productService.DeleteProduct(input.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, newDel)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "Product has been successfully deleted")
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetAllProductByCategory (c *gin.Context) {
	var input product.GetCategoryID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Category with that id not found")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newDel, err := h.productService.GetProductByCategory(input.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, newDel)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)

}