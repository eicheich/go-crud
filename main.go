package main

import (
	"go-crud/config"
	homecontroller "go-crud/controllers"
	"log"
	"net/http"
)

func main() {
	config.ConnnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("iyyah")
	http.ListenAndServe(":8080", nil)
}