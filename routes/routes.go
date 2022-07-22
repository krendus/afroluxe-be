package routes

import (
	"fmt"

	"github.com/afroluxe/afroluxe-be/config"
	"github.com/gin-gonic/gin"
)

var env config.Env = config.LoadEnv()

func SetupRoute() {
	r := gin.Default()
	CompileAuthRoute(r)
	r.Run(fmt.Sprintf(":%v", env.PORT))
}
