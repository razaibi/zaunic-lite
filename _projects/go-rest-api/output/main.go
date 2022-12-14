package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const connString = "root:adminuser@tcp(127.0.0.1:3306)/sampledb?charset=utf8mb4&parseTime=True&loc=Local"

func initializeRouter() {
	r := mux.NewRouter()

    
    r.HandleFunc("/user", GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}

