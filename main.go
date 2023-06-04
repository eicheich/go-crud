package main

import (
	"go-crud/config"
	"go-crud/controllers/categorycontroller"
	"go-crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("iyyah")
	http.ListenAndServe(":8080", nil)
}