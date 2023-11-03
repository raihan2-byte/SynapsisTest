package handler

import (
	"net/http"
	"synapsisid/helper"
	transactiondetail "synapsisid/transactionDetail"

	"github.com/gin-gonic/gin"
)

type transactionDetailsHandler struct {
	serviceTransactionDetails transactiondetail.ServiceTransactionDetails
}

func NewtransactionDetailsHandler(serviceTransactionDetails transactiondetail.ServiceTransactionDetails) *transactionDetailsHandler {
	return &transactionDetailsHandler{serviceTransactionDetails}
}

func (h *transactionDetailsHandler) GetAllTransactionDetails (c *gin.Context) {
	transactionDetails, err := h.serviceTransactionDetails.GetTransactions()
	if err != nil{
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get transactions")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, transactiondetail.FormatterGetTransaction(transactionDetails))
	c.JSON(http.StatusOK, response)
}