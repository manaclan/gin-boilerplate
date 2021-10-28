package users

import "github.com/gin-gonic/gin"

type ProfileSerializer struct {
	C *gin.Context
	User
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}
