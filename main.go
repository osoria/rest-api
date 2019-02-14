package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Car struct {
	ID    string `json:"id,omitempty"`
	Brand *Brand `json:"brand,omitempty"`
	Model string `json:"model,omitempty"`
	Power uint16 `json:"power,omitempty"`
}
type Brand struct {
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
}

var cars []Car

func GetCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cars)
}
func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}
func CreateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = params["id"]
	cars = append(cars, car)
	json.NewEncoder(w).Encode(cars)
}
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func configureRoutes(router *mux.Router) {
	router.HandleFunc("/cars", GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	router.HandleFunc("/cars/{id}", CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
}

func initData() {
	cars = append(
		cars,
		Car{ID: "1", Brand: &Brand{Name: "Dodge", Country: "United States"}, Model: "Challenger", Power: 375})
	cars = append(
		cars,
		Car{ID: "2", Brand: &Brand{Name: "Ferrari", Country: "Italy"}, Model: "GTO", Power: 400})
	cars = append(
		cars,
		Car{ID: "3", Brand: &Brand{Name: "Bugatti", Country: "France"}, Model: "Chiron", Power: 1500})
}

func main() {
	initData()

	router := mux.NewRouter()
	configureRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
