package routes

import (
	"github.com/afroluxe/afroluxe-be/controllers"
	"github.com/gin-gonic/gin"
)

func CompileAuthRoute(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/signup", controllers.HandleSignUp)
	auth.POST("/signin", controllers.HandleSignIn)
	auth.POST("/verify", controllers.VerifyEmail)
}
