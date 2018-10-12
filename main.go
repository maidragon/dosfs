package main

import (
	"log"
	// "os"
	"net/http"
	"./objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}