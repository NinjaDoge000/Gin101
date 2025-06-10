package controllers

import (
	services "gin/Services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	// register
	return &AuthController{
		authService: authService,
	}

}

func (a *AuthController) InitRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/Login", a.Login())
	auth.POST("/Register", a.Register())
}

func (a *AuthController) Login() gin.HandlerFunc {

	type user struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(c *gin.Context) {

		var userBody user

		if err := c.BindJSON(&userBody); err != nil {
			c.JSON(404, gin.H{
				"message": err,
			})

			return
		}

		myUser, err := a.authService.Login(userBody.Email, userBody.Password)

		if err != nil {
			c.JSON(400, gin.H{
				"user":    userBody,
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"user":    myUser,
			"message": "user found!",
		})
	}
}

func (a *AuthController) Register() gin.HandlerFunc {

	type user struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(c *gin.Context) {

		var userBody user

		if err := c.BindJSON(&userBody); err != nil {
			c.JSON(404, gin.H{
				"message": err,
			})

			return
		}

		myUser, err := a.authService.Register(userBody.Email, userBody.Password)

		if err != nil {
			c.JSON(404, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"user": myUser,
		})
	}
}
