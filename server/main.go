package main

import (
	"log"
)

func main() {
	serverPort := 8080
	serverInstance := CreateServer(serverPort)

	RegisterHealthService(serverInstance)

	err := serverInstance.Start()
	if err != nil {
		log.Fatal(err)
	}
}
