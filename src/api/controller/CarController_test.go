package controller

import (
	"car/infrastructure"
	"car/infrastructure/persistenceentities"
	"car/logic/operations"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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
func CreateCarResourcesWithInMemoryRepository(mockDatabaseContents map[string]persistenceentities.CarPersistenceEntity) (CarController, operations.CarOperations, infrastructure.InMemoryRepository) {
	carRepository := infrastructure.InMemoryRepository{Cars: mockDatabaseContents}
	carOperations := operations.NewCarOperations(&carRepository)
	return NewCarController(carOperations), carOperations, carRepository
}

// Test that getting all cars works
func TestGetCars(t *testing.T) {
	context, _, recorder := CreateMockEcho(
		http.MethodGet,
		"/cars",
		nil,
	)
	context.SetPath("/:vin")
	context.SetParamNames("vin")
	context.SetParamValues(persistenceentities.TestCarEntity.Vin)

	carsResource, _, _ := CreateCarResourcesWithInMemoryRepository(map[string]persistenceentities.CarPersistenceEntity{
		"JH4DB1561NS000565": persistenceentities.TestCarEntity,
		"JN8AZ2NC5B9300256": persistenceentities.TestCarEntity,
		"2FDKF38G3KCA42390": persistenceentities.TestCarEntity,
		"1GBJK39DX6E165432": persistenceentities.TestCarEntity,
	})

	if assert.NoError(t, carsResource.GetCars(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, CarsBody, recorder.Body.String())
	}
}
