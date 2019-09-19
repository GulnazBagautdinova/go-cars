package main

import (
	"github.com/gorilla/mux"
	//"go-cars/app"
	"os"
	"fmt"
	"net/http"
	"go-cars/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/models", controllers.createModel).Methods("POST")
	router.HandleFunc("/api/models", controllers.getModels).Methods("GET")
	router.HandleFunc("/api/models/{id}", controllers.getModel).Methods("GET")
	router.HandleFunc("/api/models/{id}", controllers.deleteModel).Methods("DELETE")
	router.HandleFunc("/api/me/contacts", controllers.updateModel).Methods("PUT") 

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}