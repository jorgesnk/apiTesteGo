package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Person struct {
	ID        string  `json:"id,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Address   Address `json:"address,omitempty"`
}

func main() {
	fmt.Print("teste")
	router := mux.NewRouter()

	people = append(people, Person{
		Address:   Address{City: "teste", State: "teste"},
		Firstname: "teste",
		ID:        "1",
		Lastname:  "teste",
	})

	router.HandleFunc("/people", GetPeople).Methods("GET")

	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")

	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")

	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	router.HandleFunc("/people/{id}", Update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))

}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(people)

}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {

		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var personUpdate Person
	_ = json.NewDecoder(r.Body).Decode(&personUpdate)

	println(personUpdate.ID, "teste")

	for i, item := range people {
		if item.ID == params["id"] {
			people[i] = personUpdate
			break
		}
	}

	json.NewEncoder(w).Encode(personUpdate)

}
