package main

import (
	"drinklytics/internal/config"
	"drinklytics/internal/controllers"
	"drinklytics/internal/middleware"
	router "drinklytics/internal/routers"
	"drinklytics/internal/services"
	"log"
	"time"

	"drinklytics/internal/db"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	log.Println("Application Initiated")
	go middleware.CleanClients()
	loc, _ := time.LoadLocation(config.GinTZ)
	time.Local = loc
	gin.SetMode(config.GinMode)
	db, err := db.InitDb()
	if err != nil {
		log.Printf("Error Initializing DB: %v \n", err)
	}

	authService := services.NewAuthService(db)
	userService := services.NewUserService(db)
	finService := services.NewFinService(db)

	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	finController := controllers.NewFinController(finService)

	router.InitRoutes(authController, userController, finController)
}
