package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("The API is running on port", config.Port)

	r := router.Create()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}