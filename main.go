package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("The API is running!")

	r := router.Create()

	log.Fatal(http.ListenAndServe(":5000", r))
}