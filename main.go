package main

import (
	"go-crud/config"
	"log"
	"net/http"
)

func main() {
	config.ConnnectDB()

	log.Println("iyyah")
	http.ListenAndServe(":8080", nil)
}