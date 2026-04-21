package router

import (
	"fmt"

	// "html/template"

	"drinklytics/internal/config"
	"drinklytics/internal/controllers"
	"drinklytics/internal/middleware"

	"github.com/flosch/pongo2/v6"

	"github.com/gin-gonic/gin"
)



func InitRoutes(authController *controllers.AuthController, userController *controllers.UserController, finController *controllers.FinController) {
	router := gin.Default()
	router.Use(middleware.RateLimiter())
	// router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "./web/static")
	// router.GET("/", func(c *gin.Context) {
	// 	middleware.Render(c, "index.html", pongo2.Context{
	// 		"name":   "Login",
	// 	})
	// })

	router.GET("/getin", func(c *gin.Context) {
		successMsg := c.Query("success")
		if successMsg == "" {
			successMsg = ""
		}
		errorMsg := c.Query("error")
		if errorMsg == "" {
			errorMsg = ""
		}
		middleware.Render(c, "temp_login.html", pongo2.Context{
			"title":   "Login",
			"success": successMsg,
			"error": errorMsg,
		})
	})

	router.POST("/getin", userController.CreateUser)

	router.GET("/login", func(c *gin.Context) {
		successMsg := c.Query("success")
		if successMsg == "" {
			successMsg = ""
		}
		middleware.Render(c, "login.html", pongo2.Context{
			"title":   "Login",
			"success": successMsg,
		})
	})
	router.GET("/signup", func(c *gin.Context) {
		errorMsg := c.Query("error")
		if errorMsg == "" {
			errorMsg = ""
		}
		middleware.Render(c, "signup.html", pongo2.Context{
			"title": "Sign Up",
			"error": errorMsg,
		})
	})

	app := router.Group("/drinklytics/v1")
	app.Static("/static", "./web/static")

	router.GET("/", func(c *gin.Context) {
		middleware.Render(c, "index.html", pongo2.Context{
			"name":   "Login",
		})
	})
	// app := router.Group("/drinklytics/v1")
	// app.POST("/login", authController.LoginUser)          // can be accessed by Admin, Analyst and Viewer
	// app.POST("/register-user", userController.CreateUser) // can be accessed by Admin, Analyst and Viewer

	// app.Use(middleware.TokenValidator)
	// app.POST("/logout", authController.LogoutUser) // can be accessed by Admin, Analyst and Viewer

	// users := app.Group("/user")
	// users.PATCH("/update", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), userController.UpdateUser)                  // can be accessed by Admin, Analyst and Viewer
	// users.PATCH("/update-password", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), userController.UpdateUserPassword) // can be accessed by Admin, Analyst and Viewer
	// users.DELETE("/delete", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), userController.DeleteUser)                 // can be accessed by Admin, Analyst and Viewer
	// users.PATCH("/update-role", middleware.Authorize(config.Admin), userController.UpdateUserRole)                                        // can be accessed by Admin
	// users.PATCH("/update-status", middleware.Authorize(config.Admin), userController.UpdateUserStatus)                                    // can be accessed by Admin
	// users.GET("/get/:id", middleware.Authorize(config.Admin), userController.GetUser)                                                     // can be accessed by Admin
	// users.GET("/get-all", middleware.Authorize(config.Admin), userController.GetAllUsers)                                                 // can be accessed by Admin
	// users.DELETE("/hard-delete", middleware.Authorize(config.Admin), userController.HardDeleteUser)                                       // can be accessed by Admin

	// transactions := app.Group("/transactions")
	// transactions.GET("/transaction-type", middleware.Authorize(config.Admin, config.Analyst), finController.GetAllTransactionTypes) // can be accessed by Admin, Analyst
	// transactions.POST("/transaction-type", middleware.Authorize(config.Admin), finController.CreateTransactionTypes)                // can be accessed by Admin
	// transactions.PUT("/transaction-type", middleware.Authorize(config.Admin), finController.UpdateTransactionTypes)                 // can be accessed by Admin
	// transactions.DELETE("/transaction-type/:id", middleware.Authorize(config.Admin), finController.DeleteTransactionTypes)          // can be accessed by Admin

	// transactions.GET("/trx", middleware.Authorize(config.Admin, config.Analyst), finController.GetTransaction) // can be accessed by Admin, Analyst
	// transactions.POST("/trx", middleware.Authorize(config.Admin), finController.CreateTransaction)             // can be accessed by Admin,
	// transactions.PUT("/trx", middleware.Authorize(config.Admin), finController.UpdateTransaction)              // can be accessed by Admin
	// transactions.DELETE("/trx/:id", middleware.Authorize(config.Admin), finController.DeleteTransaction)       // can be accessed by Admin

	// dashboard := app.Group("/dashboard")
	// dashboard.GET("/financial-summary", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), finController.FinancialSummary) // can be accessed by Admin, Analyst and Viewer
	// dashboard.GET("/category-summary", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), finController.CategorySummary)   // can be accessed by Admin, Analyst and Viewer
	// dashboard.GET("/recent-activity", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), finController.RecentActivity)     // can be accessed by Admin, Analyst and Viewer
	// dashboard.GET("/trends", middleware.Authorize(config.Admin, config.Analyst, config.Viewer), finController.Trends)                      // can be accessed by Admin, Analyst and Viewer

	router.Run(fmt.Sprintf(`:%s`, config.AppPort))
}
