package routes

import (
	"fmt"

	"github.com/afroluxe/afroluxe-be/config"
	"github.com/afroluxe/afroluxe-be/controllers"
	"github.com/gin-gonic/gin"
)

var env config.Env = config.LoadEnv()

func SetupRoute() {
	r := gin.Default()
	r.GET("/", controllers.WelcomeHandler)
	CompileAuthRoute(r)
	r.Run(fmt.Sprintf(":%v", env.PORT))
}
