package controller

import (
	"car/infrastructure/connectedcars"
	persistenceentities2 "car/infrastructure/connectedcars/entities"
	"car/logic/operations"
)

// A valid request body for a Car
var CarBodyRequest = "{\"Vin\":\"JH4DB1561NS000565\",\"Brand\":\"VW\",\"Model\":\"ID2\"}\n"

// A valid response body for a Car
var CarBodyResponse = "{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"}\n"

// A valid response body for Cars
var CarsBody = "{\"Cars\":[{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"},{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"},{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"},{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"}]}\n"

// List of invalid Vins according to the domain constraints
var InvalidVins = []string{
	"JH4DA3350KS00",
	"2CIGP44362R700796",
	"1C3CDZBG8DN5O4146",
	"1gCDC14K2LE198114",
	"1G3NF52E3XC4036521",
}

// Create all resources used by the car controller with an underlying in-memory repository
func CreateCarResourcesWithInMemoryRepository(mockDatabaseContents []persistenceentities2.ConnectedCarsEntity) (CarController, operations.CarOperations, connectedcars.ConnectedCars) {
	carRepository := connectedcars.ConnectedCars{Cars: mockDatabaseContents}
	carOperations := operations.NewCarOperations(&carRepository)
	return NewCarController(carOperations), carOperations, carRepository
}
