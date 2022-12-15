package router

import (
	"github.com/gorilla/mux"
	"github.com/kvbendalam/cars/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/cars", middleware.ReadAllCars).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/car/{id}", middleware.RedCar).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cars", middleware.CreateCar).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/car/{id}", middleware.DeleteCar).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/car/{id}", middleware.UpdateCar).Methods("PUT", "OPTIONS")
	return router
}
