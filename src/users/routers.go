package users

import (
	"github.com/gin-gonic/gin"
	"github.com/manaclan/gin-boilerplate/src/database"
)

type UsersRouter struct {
	Services    Services
	Controllers Controllers
}

func (router UsersRouter) Init() {
	router.Services = Services{Client: database.DB}
	router.Controllers = Controllers{Services: router.Services}
}

func (ur UsersRouter) Route(router *gin.RouterGroup) {
	router.POST("/login", ur.Controllers.Login)
	router.POST("/register", ur.Controllers.Register)
}
