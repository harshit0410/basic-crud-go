package main

import (
	"JobWorker/database"
	"JobWorker/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
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

	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/create", createPerson).Methods("POST")
	router.HandleFunc("/get/{id}", getPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", updatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", deletPersonByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person model.Person
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person model.Person
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func updatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person model.Person
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func deletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person model.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
