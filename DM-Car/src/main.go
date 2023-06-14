package main

import (
	"car/DM-Car/src/api/controller"
	"car/DM-Car/src/infrastructure"
	"car/DM-Car/src/logic/operations"
	"log"
	"net/http"
)

func main() {
	carOperations := operations.NewCarOperations(infrastructure.NewCarRepository())
	carCollectionResource := controller.NewCarCollectionResource(carOperations)

	http.HandleFunc("/car", carCollectionResource.HandleAddCar)
	http.HandleFunc("/car/", carCollectionResource.HandleGetCar)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("System failed to listen and serve car collection resource: %v", err)
	}
}
