package handler

import (
	"net/http"
	"strconv"
	"synapsisid/category"
	"synapsisid/helper"
	"synapsisid/user"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.ServiceCategory
}

func NewCategoryHandler(service category.ServiceCategory) *categoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.categoryService.CreateCategory(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := category.FormatterCategory(newUser)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdatedCategory(c *gin.Context) {
	var inputID category.GetinputCategoryID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData category.UpdatedCategory

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.categoryService.UpdatedCategory(inputID, inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := category.FormatterCategoryUpdated(newUser)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)

}

func (h *categoryHandler) DeletedCategory(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	userID := currentUser.ID

	newDel, err := h.categoryService.DeleteCategory(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, newDel)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// responseDeleted := "Your account has been successfully deleted"

	response := helper.APIresponse(http.StatusOK, "Category has been successfully deleted")
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetAllCategory (c *gin.Context){
	Categorys, err := h.categoryService.GetCategorys()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get Category")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, Categorys)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategory(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	category, err := h.categoryService.GetCategory(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get category")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, category)
	c.JSON(http.StatusOK, response)
}
