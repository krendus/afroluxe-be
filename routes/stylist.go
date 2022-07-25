package routes

import (
	"github.com/afroluxe/afroluxe-be/controllers"
	"github.com/gin-gonic/gin"
)

func CompileStylistRoute(r *gin.Engine) {
	auth := r.Group("/stylist")
	auth.GET("/:id", controllers.GetStylist)
}
