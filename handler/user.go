package handler

import (
	"fmt"
	"net/http"
	"synapsisid/auth"
	"synapsisid/helper"
	"synapsisid/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
    var input user.RegisterUserInput

    err := c.ShouldBindJSON(&input)
    if err != nil {
        errors := helper.FormatValidationError(err)
        errorMessage := gin.H{"errors": errors}
        response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

    // Periksa ketersediaan email sebelum mendaftarkan pengguna
    isEmailAvailable, err := h.userService.IsEmaillAvailabilty(input.Email)
    if err != nil {
        response := helper.APIresponse(http.StatusInternalServerError, nil)
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    // Jika email tidak tersedia, kirim respons kesalahan
    if !isEmailAvailable {
        response := helper.APIresponse(http.StatusConflict, nil)
        c.JSON(http.StatusConflict, response)
        return
    }

    // Register user jika email tersedia
    newUser, err := h.userService.RegisterUser(input)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

    // Format dan kirim respons berhasil jika registrasi berhasil
    formatter := user.FormatterRegister(newUser)
    response := helper.APIresponse(http.StatusOK, formatter)
    c.JSON(http.StatusOK, response)
}


func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID, loggedinUser.Role)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatterUser(token)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdatedUser(c *gin.Context) {
	
	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	inputData  := currentUser.ID


	var inputBalance user.UpdatedUser

	err := c.ShouldBindJSON(&inputBalance)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	

	newUser, err := h.userService.UpdatedUser(inputData, inputBalance)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := fmt.Sprintf("Your balance has been succesfully updated to Rp %d", newUser.Balance)

	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)

}

// func (h *userHandler) DeletedUser(c *gin.Context) {
// 	// var input user.DeletedUser

// 	currentUser := c.MustGet("currentUser").(user.User)
// 	//ini inisiasi userID yang mana ingin mendapatkan id si user
// 	userID := currentUser.ID

// 	// err := c.ShouldBindUri(&input)
// 	// if err != nil {
// 	// 	errors := helper.FormatValidationError(err)
// 	// 	errorMessage := gin.H{"errors": errors}
// 	// 	response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
// 	// 	c.JSON(http.StatusUnprocessableEntity, response)
// 	// 	return
// 	// }

// 	newDel, err := h.userService.DeleteUser(userID)
// 	if err != nil {
// 		response := helper.APIresponse(http.StatusUnprocessableEntity, newDel)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	// responseDeleted := "Your account has been successfully deleted"

// 	response := helper.APIresponse(http.StatusOK, "Your account has been successfully deleted")
// 	c.JSON(http.StatusOK, response)
// }
