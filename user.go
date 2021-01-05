package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	//_"github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB


func AllUsers(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "All users Endpoint hit")
}


func NewUser(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "New user Endpoint hit")
}


func DeleteUser(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "Delete user Endpoint hit")
}


func UpdateUser(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "Update user Endpoint hit")
}





