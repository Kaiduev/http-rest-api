package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil{
		panic("Could not connect to the databse")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "New Users Endpoint hit")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Users Endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update Users Endpoint hit")
}
