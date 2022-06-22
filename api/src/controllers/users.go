package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.UserRepositoryFactory(db)

	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Inserted id: %d", userId)))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("listing users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting an user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating an user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting an user"))
}
