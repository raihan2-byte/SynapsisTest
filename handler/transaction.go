package handler

import (
	"net/http"
	"strconv"
	"synapsisid/helper"
	"synapsisid/transaction"
	"synapsisid/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.ServiceTransaction
}

func NewtransactionHandler(service transaction.ServiceTransaction) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var inputID transaction.GetIDProduct

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input transaction.TransactionInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	inputData  := currentUser.ID

	newUser, err := h.transactionService.CreateTransaction(inputID, inputData, input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.transactionService.SaveToTransaction(inputData, newUser.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, transaction.FormatterGet(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransaction(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("product_id"))
	userID, _ := strconv.Atoi(c.Query("user_id"))

	sosmed, err := h.transactionService.GetTransaction(productID, userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get transaction")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, transaction.FormatterGetTransaction(sosmed))
	c.JSON(http.StatusOK, response)
}
