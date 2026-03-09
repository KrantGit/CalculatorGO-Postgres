package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/calculator", calculatorHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
