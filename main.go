package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kvbendalam/cars/router"
)

func main() {
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router.Router()))

}
