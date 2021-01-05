package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	)

func helloWorld(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("authentification Microservice"))
}

func handleRequest(){
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":8001",myRouter))
}

func main(){
	fmt.Println("authentification Microservice")

}