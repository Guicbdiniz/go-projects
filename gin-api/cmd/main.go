package main

import (
	"fmt"
	"log"

	"github.com/Guicbdiniz/go-projects/gin-api/pkg/api"
)

const port = 3001

func main() {
	api, err := api.CreateAPI()

	if err != nil {
		log.Fatalf("Error captured while creating the API, %v", err)
	}

	portString := fmt.Sprintf(":%d", port)
	log.Fatalln(api.Run(portString))
}
