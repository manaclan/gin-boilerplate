package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manaclan/gin-boilerplate/src/database"
	"github.com/manaclan/gin-boilerplate/src/users"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	database.Init()
	r := gin.Default()

	usersRouteGroup := r.Group("/users")
	usersRouter := users.UsersRouter{}
	usersRouter.Init()
	usersRouter.Route(usersRouteGroup)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Run(":8525")
}
