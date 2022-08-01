package main

import (
	"fmt"

	"github.com/afroluxe/afroluxe-be/routes"
	"github.com/afroluxe/afroluxe-be/utils"
)

func main() {
	fmt.Println(utils.GenerateRandomOtp(6))
	routes.SetupRoute()
}
