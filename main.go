package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("authentification Microservice"))
}


func handleRequest(){
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	port := ":8001"
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func main(){
	fmt.Println("authentification Microservice")
	handleRequest()
}