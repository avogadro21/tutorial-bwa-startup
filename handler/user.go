package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutorial_bwa_startup/global"
	"tutorial_bwa_startup/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		message := "Error while registering account"
		errors := global.ValidationError(err)
		errorMessage := gin.H{"errors": errors}

		errorResponse := global.APIResponse("error", http.StatusUnprocessableEntity, message, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		message := "Error while registering account"
		errorResponse := global.APIResponse("error", http.StatusBadRequest, message, nil)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	formatter := user.GetDTORegisterUser(newUser)
	response := global.APIResponse("success", http.StatusCreated, "Account has been created", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		message := "Login Error"
		errors := global.ValidationError(err)
		errorMessage := gin.H{"errors": errors}

		errorResponse := global.APIResponse("error", http.StatusBadRequest, message, errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	newUser, err := h.userService.Login(input)
	if err != nil {
		message := "Login Error"
		errorMessage := gin.H{"errors": err.Error()}
		errorResponse := global.APIResponse("error", http.StatusBadRequest, message, errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	formatter := user.GetDTOLogin(newUser, "token")
	response := global.APIResponse("success", http.StatusOK, "Login Success", formatter)
	c.JSON(http.StatusOK, response)
}
