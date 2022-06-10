package main

import (
	"JobWorker/controller"
	"JobWorker/database"
	"JobWorker/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	initDB()
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	initRoute(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "12345",
			DB:         "job_scheduler",
		}

	err := database.Connect(config)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&model.Person{})
}

func initRoute(router *mux.Router) {

	router.HandleFunc("/get", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/create", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", controller.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controller.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.DeletPersonByID).Methods("DELETE")
}
