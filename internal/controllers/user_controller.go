package controllers

import (
	// "drinklytics/internal/config"
	// model "drinklytics/internal/models"
	"drinklytics/internal/services"
	"fmt"
	// "log"
	// "net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"drinklytics/internal/middleware"
	"github.com/flosch/pongo2/v6"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) CreateUser(c *gin.Context) {
	fmt.Println("dats ", c.PostForm("username"))
	// if c.PostForm("username") == "" {
	// 	fmt.Println("errror ")
	// 	middleware.Render(c, "temp_login.html", pongo2.Context{
	// 		"title":   "Login",
	// 		"error": "Please enter username",
	// 	})
	// 	// c.Redirect(http.StatusSeeOther, "/getin?error=User+already+exists")
	// }

	// user, err := h.userService.CreateUser(c.PostForm("username"))
	// if err != nil {
	// 	fmt.Println("eror ", err)
	// 	middleware.Render(c, "temp_login.html", pongo2.Context{
	// 		"title":   "Login",
	// 		"error": "User already exists",
	// 	})
	// 	// c.Redirect(http.StatusSeeOther, "/getin?error=User+already+exists")s

	// } else {
		middleware.Render(c, "index.html", pongo2.Context{
			// "user":   user,
			"user":   c.PostForm("username"),
		})
	// }
}


// func (h *UserController) CreateUser(c *gin.Context) {
// 	// var request model.UserRequests
// 	fmt.Println("dats ", c.PostForm("username"),c.PostForm("email"),c.PostForm("password"))
// 	// if err := c.ShouldBindJSON(&request); err != nil {
// 	// 	log.Println("Error Binding Json", err)
// 	// 	// c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	// 	c.Redirect(http.StatusBadRequest, "/signup?error=All+fields+required")
// 	// 	return
// 	// }

// 	// _, err := h.userService.CreateUser(&request) //result
// 	// if err != nil {
// 		// c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		c.Redirect(http.StatusSeeOther, "/getin?error=User+already+exists")

// 	// } else {
// 	// 	// c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 		// c.Redirect(http.StatusSeeOther, "/login?success=Account+created")
// 	// }
// }

// func (h *UserController) DeleteUser(c *gin.Context) {
// 	var request model.UserRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Println("Error Binding Json", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		return
// 	}

// 	role := c.GetString("role")
// 	if role != string(config.Admin) {
// 		userId, _ := c.Get("userId")
// 		request.ID = userId.(int64)
// 	}

// 	err := h.userService.DeleteUser(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", "Successfully Deleted", nil))
// 	}
// }

// func (h *UserController) HardDeleteUser(c *gin.Context) {
// 	var request model.UserRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Println("Error Binding Json", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		return
// 	}

// 	role := c.GetString("role")
// 	if role != string(config.Admin) {
// 		userId, _ := c.Get("userId")
// 		request.ID = userId.(int64)
// 	}

// 	err := h.userService.HardDeleteUser(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", "Successfully Deleted", nil))
// 	}
// }

// func (h *UserController) UpdateUser(c *gin.Context) {
// 	var request model.UserRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Println("Error Binding Json", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		return
// 	}

// 	role := c.GetString("role")
// 	if role != string(config.Admin) {
// 		userId, _ := c.Get("userId")
// 		request.ID = userId.(int64)
// 	}

// 	result, err := h.userService.UpdateUser(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 	}
// }

// func (h *UserController) UpdateUserPassword(c *gin.Context) {
// 	var request model.UserRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Println("Error Binding Json", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		return
// 	}

// 	role := c.GetString("role")
// 	if role != string(config.Admin) {
// 		userId, _ := c.Get("userId")
// 		request.ID = userId.(int64)
// 	}

// 	result, err := h.userService.UpdateUserPassword(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 	}
// }

// func (h *UserController) UpdateUserRole(c *gin.Context) {
// 	var request model.User
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Println("Error Binding Json", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		return
// 	}

// 	result, err := h.userService.UpdateUserRole(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 	}
// }

// func (h *UserController) UpdateUserStatus(c *gin.Context) {
// 	var request model.User
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Println("Error Binding Json", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 		return
// 	}

// 	result, err := h.userService.UpdateUserStatus(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 	}
// }

// func (h *UserController) GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var userId int64 = -1
// 	var err error
// 	if id != "" {
// 		userId, err = strconv.ParseInt(id, 10, 64)
// 	}
// 	if err != nil {
// 		log.Println("Error Parsing int to string", err)
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	}

// 	result, err := h.userService.GetUser(userId)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 	}
// }

// func (h *UserController) GetAllUsers(c *gin.Context) {
// 	result, err := h.userService.GetAllUsers()

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.BuildErrorResponse("Bad Request", nil, err))
// 	} else {
// 		c.JSON(http.StatusOK, model.BuildResponse("Message", result, nil))
// 	}
// }
