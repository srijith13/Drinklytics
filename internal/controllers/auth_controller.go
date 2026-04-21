package controllers

import (
	model "drinklytics/internal/models"
	"drinklytics/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (h *AuthController) LoginUser(c *gin.Context) {
	var request model.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error Binding Json", err)
		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
		return
	}

	// // Custom validator to check the request body matches the required format and data structure to prevent injections through json
	// if errors := helper.ValidateCustomBody(request); len(errors) != 0 {
	// 	c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", errors, errors))
	// 	return
	// }

	result, err := h.authService.LoginUser(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", result, err))
	} else {
		c.JSON(http.StatusOK, model.BuildResponse("Message", result, err))
	}
}

func (h *AuthController) LogoutUser(c *gin.Context) {
	userId, _ := c.Get("userId")
	err := h.authService.LogoutUser(userId.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
	} else {
		c.JSON(http.StatusOK, model.BuildResponse("Message", "Logged Out", nil))
	}
}
