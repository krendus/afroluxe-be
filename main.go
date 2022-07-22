package main

import (
	"fmt"

	"github.com/afroluxe/afroluxe-be/routes"
)

func main() {
	routes.SetupRoute()
	fmt.Println("Server started...")
}
