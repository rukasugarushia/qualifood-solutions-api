package main

import (
	"log"
	"qualifood-solutions-api/internal/infrastructure"
)

func main() {
	router := infrastructure.SetupRouter()
	log.Println("Server is running at http://localhost:8080")
	router.Run(":8080")
}
