package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kvbendalam/cars/models"
)

var carsDB []models.Car
var carIDTrack = 0

func init() {
	fmt.Println("Intializing data")
	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "Nexon", Price: 1500000})
	carIDTrack++
	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "Harrier", Price: 2500000})
}

func ReadAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := readAllCars()
	json.NewEncoder(w).Encode(payload)
}

func readAllCars() []models.Car {
	var cars []models.Car
	for _, element := range carsDB {
		cars = append(cars, element)
	}
	return cars
}

func RedCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	payload := redCar(vars["id"])

	json.NewEncoder(w).Encode(payload)
}

func redCar(id string) models.Car {
	var car models.Car

	for _, element := range carsDB {
		if element.ID == id {
			car = element
			break
		}
	}
	return car
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var car models.Car

	_ = json.NewDecoder(r.Body).Decode(&car)
	payload := createCar(car)
	json.NewEncoder(w).Encode(payload)
}

func createCar(car models.Car) models.Car {
	car.ID = strconv.Itoa(carIDTrack)
	carIDTrack++
	carsDB = append(carsDB, car)
	return car
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	payload := deleteCar(vars["id"])
	json.NewEncoder(w).Encode(payload)
}

func deleteCar(id string) models.Car {
	var car models.Car
	for index, element := range carsDB {
		if element.ID == id {
			carsDB = append(carsDB[:index], carsDB[index+1:]...)
			car = element
			break
		}
	}
	return car
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	var car models.Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	payload := updateCar(vars["id"], car)
	json.NewEncoder(w).Encode(payload)
}

func updateCar(id string, car models.Car) models.Car {
	for index, element := range carsDB {
		if element.ID == id {
			element.Model = car.Model
			element.Price = car.Price
			carsDB[index] = element
			break
		}
	}
	return car
}
