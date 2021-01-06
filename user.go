package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
}

func InitialMigration(){

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err.Error())
		panic("Migrations failed")
	}
}

func AllUsers(w http.ResponseWriter, req *http.Request){
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Could not connect to database")
	}
	//defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}


func NewUser(w http.ResponseWriter, req *http.Request){
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Could not connect to database")
	}

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{
		Email: email,
		Name: name,
	})

	fmt.Fprint(w, "New user Successfully Created")
}


func DeleteUser(w http.ResponseWriter, req *http.Request){
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Could not connect to database")
	}

	vars := mux.Vars(req)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprint(w, "User Successfully Deleted")

}


func UpdateUser(w http.ResponseWriter, req *http.Request){
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Could not connect to database")
	}

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)

	fmt.Fprint(w, "User Successfully Updated")
}



