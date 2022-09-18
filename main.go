package main

// @title           Afroluxe API
// @version         1.0
// @description     This is the API documentation of afroluxe.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      afroluxe.herokuapp.com
// @BasePath  /
// @schemes  https

// @securityDefinitions.basic  BasicAuth

import (
	_ "github.com/afroluxe/afroluxe-be/docs"
	"github.com/afroluxe/afroluxe-be/routes"
)

func main() {
	routes.SetupRoute()
}
