package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("iyyah")
	http.ListenAndServe(":8080", nil)
}