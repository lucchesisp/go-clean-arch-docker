package main

import (
	"fmt"
	"github.com/lucchesisp/go-clean-arch-docker/src/config"
	"github.com/lucchesisp/go-clean-arch-docker/src/routes"
	"log"
)

func main() {
	// Load enviroments variables
	port, err := config.GetEnvVariable("PORT")

	if err != nil {
		log.Fatal("Error getting port from env variables")
		return
	}

	fmt.Printf("Starting server on port %s\n", port)
	routes.Run(port)
}
