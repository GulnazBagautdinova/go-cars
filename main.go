package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"github.com/GulnazBagautdinova/go-cars/controllers"
)


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/models", controllers.CreateModel).Methods("POST")
	router.HandleFunc("/api/models/{id}", controllers.UpdateModel).Methods("PUT")
	router.HandleFunc("/api/models/{id}", controllers.GetModel).Methods("GET")
	router.HandleFunc("/api/models/{id}", controllers.DeleteModel).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) 
	if err != nil {
		fmt.Print(err)
	}
}