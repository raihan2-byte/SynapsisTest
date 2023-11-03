package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"synapsisid/cart"
	"synapsisid/helper"
	transactiondetail "synapsisid/transactionDetail"
	"synapsisid/user"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	cartService cart.ServiceCarts
}

func NewAddToCartHandler(cartService cart.ServiceCarts) *cartHandler {
	return &cartHandler{cartService}
}

func (h *cartHandler) AddToCart (c *gin.Context) {
	var productId cart.GetIdProduct

	err := c.ShouldBindUri(&productId)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input cart.InputCart

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

	newUser, err := h.cartService.AddToCart(productId, inputData, input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, cart.FormatterCartPost(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) UpdateCart(c *gin.Context) {
    cartID, err := strconv.Atoi(c.Param("cartID"))
    if err != nil {
        response := helper.APIresponse(http.StatusBadRequest, "Invalid cart ID")
        c.JSON(http.StatusBadRequest, response)
        return
    }

    var input cart.InputCart
    err = c.ShouldBindJSON(&input)
	if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, "Invalid body input")
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	inputData  := currentUser.ID

	// userID := c.MustGet("currentUser").(int)


    updatedCart, err := h.cartService.UpdatedCart(cartID, input, inputData)
    if err != nil {
        response := helper.APIresponse(http.StatusInternalServerError, "Failed to update cart")
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    response := helper.APIresponse(http.StatusOK, cart.FormatterCartPost(updatedCart))
    c.JSON(http.StatusOK, response)
}


func (h *cartHandler) GetAllChart(c *gin.Context) {
	
	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	inputData  := currentUser.ID

	//seperti biasa untuk berikan respon dan gunakan userID yang uda di inisiasi
	carts, err := h.cartService.GetAllCartByUserId(inputData)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "failed get all cart")
        c.JSON(http.StatusUnprocessableEntity, response)
        return
	}
	response := helper.APIresponse(http.StatusOK, cart.FormatterGetCart(carts))
    c.JSON(http.StatusOK, response)

}

func (h *cartHandler) PayByUserID(c *gin.Context) {
	userID := c.MustGet("currentUser").(user.User).ID

	// Panggil service untuk melakukan pembayaran berdasarkan userID
	transactionDetails, err := h.cartService.PayByUserID(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to process payment: "+err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, transactionDetails)
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) PayByCartID (c *gin.Context) {
	var IdCart cart.GetIDCart

	err := c.ShouldBindUri(&IdCart)
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

	transactionDetails, err := h.cartService.PayFromCartID(IdCart, inputData)
	fmt.Println(IdCart)
	fmt.Println(inputData)
	if err != nil {
        response := helper.APIresponse(http.StatusBadRequest, "Invalid cart ID")
        c.JSON(http.StatusBadRequest, response)
        return
    }

	response := helper.APIresponse(http.StatusOK, transactiondetail.FormatterGet(transactionDetails))
	c.JSON(http.StatusOK, response)
} 

func (h *cartHandler) DeleteCart (c *gin.Context) {
	var inputID cart.GetIDCart
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Invalid cart ID")
        c.JSON(http.StatusBadRequest, response)
        return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	inputData  := currentUser.ID

	_, err = h.cartService.DeleteCart(inputData, inputID)
	// fmt.Println(cartID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Invalid cart ID")
        c.JSON(http.StatusBadRequest, response)
        return
	}
	response := helper.APIresponse(http.StatusOK, "Cart has been successfully deleted")
	c.JSON(http.StatusOK, response)

}