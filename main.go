package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	controllers "github.com/manaclan/gin-boilerplate/src/1-controllers"
	services "github.com/manaclan/gin-boilerplate/src/2-services"
	"github.com/manaclan/gin-boilerplate/src/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	client := database.ConnectDatabase()
	servc := services.Services{Client: client}
	ctrler := controllers.Controllers{Services: servc}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.POST("/login", ctrler.Login)
	r.POST("/register", ctrler.Register)
	r.Run(":8525")
}
