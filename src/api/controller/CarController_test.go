package controller

import (
	"car/src/infrastructure"
	entities2 "car/src/infrastructure/entities"
	"car/src/logic/operations"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A valid request body for a Car
var CarBodyRequest = "{\"Vin\":\"JH4DB1561NS000565\",\"Brand\":\"VW\",\"Model\":\"ID2\"}\n"

// A valid response body for a Car
var CarBodyResponse = "{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"}\n"

// A valid response body for Cars
var CarsBody = "{\"Cars\":[{\"Vin\":{\"Vin\":\"JH4DB1561NS000565\"},\"Brand\":\"VW\",\"Model\":\"ID2\"},{\"Vin\":{\"Vin\":\"JN8AZ2NC5B9300256\"},\"Brand\":\"VW\",\"Model\":\"ID2\"},{\"Vin\":{\"Vin\":\"2FDKF38G3KCA42390\"},\"Brand\":\"VW\",\"Model\":\"ID2\"},{\"Vin\":{\"Vin\":\"1GBJK39DX6E165432\"},\"Brand\":\"VW\",\"Model\":\"ID2\"}]}\n"

// List of invalid Vins according to the domain constraints
var InvalidVins = []string{
	"JH4DA3350KS00",
	"2CIGP44362R700796",
	"1C3CDZBG8DN5O4146",
	"1gCDC14K2LE198114",
	"1G3NF52E3XC4036521",
}

// Create all resources used by the car controller with an underlying mocked repository
func CreateCarResourcesWithMockRepository(mockDatabaseContents map[string]entities2.CarPersistenceEntity) (CarController, operations.CarOperations, infrastructure.MockCarRepository) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: mockDatabaseContents}
	carOperations := operations.NewCarOperations(&carRepository)
	return NewCarController(carOperations), carOperations, carRepository
}

// Test that adding a car works
func TestAddCar(t *testing.T) {
	context, request, recorder := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(CarBodyRequest),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, carRepository := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{})

	if assert.NoError(t, carsResource.AddCar(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, carRepository.MockDatabase, entities2.TestCarEntity.Vin.Vin)
		assert.Equal(t, carRepository.MockDatabase[entities2.TestCarEntity.Vin.Vin], entities2.TestCarEntity)
	}
}

// Test that adding a car with an existing Vin does not work
func TestAddCarWithExistingVin(t *testing.T) {
	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(CarBodyRequest),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{
		"JH4DB1561NS000565": entities2.TestCarEntity,
	})

	assert.Error(t, carsResource.AddCar(context))
}

// Test that adding a car with an invalid Vin does not work
func TestAddCarInvalidVin(t *testing.T) {
	for _, invalidVin := range InvalidVins {
		body := fmt.Sprintf(`
		{
			"vin": "%s",
			"brand": "%s",
			"model": "%s"
		}
		`,
			invalidVin,
			entities2.TestCarEntity.Brand,
			entities2.TestCarEntity.Model)

		context, request, _ := CreateMockEcho(
			http.MethodPost,
			"/cars",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{})

		assert.Error(t, carsResource.AddCar(context))
	}
}

// Test that adding a car without a Vin does not work
func TestAddCarNoVin(t *testing.T) {
	body := fmt.Sprintf(`
	{
		"brand": "%s",
		"model": "%s"
	}
	`,
		entities2.TestCarEntity.Brand,
		entities2.TestCarEntity.Model)

	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{})

	assert.Error(t, carsResource.AddCar(context))
}

// Test that adding a car without a brand does not work
func TestAddCarNoBrand(t *testing.T) {
	body := fmt.Sprintf(`
	{
		"vin": "%s",
		"model": "%s"
	}
	`,
		entities2.TestCarEntity.Vin.Vin,
		entities2.TestCarEntity.Model)

	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{})

	assert.Error(t, carsResource.AddCar(context))
}

// Test that adding a car without a model does not work
func TestAddCarNoModel(t *testing.T) {
	body := fmt.Sprintf(`
	{
		"vin": "%s",
		"brand": "%s"
	}
	`,
		entities2.TestCarEntity.Vin.Vin,
		entities2.TestCarEntity.Brand)

	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{})

	assert.Error(t, carsResource.AddCar(context))
}

// Test that getting a specific car works
func TestGetCar(t *testing.T) {
	context, _, recorder := CreateMockEcho(
		http.MethodGet,
		"/cars",
		nil,
	)
	context.SetPath("/:vin")
	context.SetParamNames("vin")
	context.SetParamValues(entities2.TestCarEntity.Vin.Vin)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{
		"JH4DB1561NS000565": entities2.TestCarEntity,
	})

	if assert.NoError(t, carsResource.GetCar(context, entities2.TestCarEntity.Vin.Vin)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, CarBodyResponse, recorder.Body.String())
	}
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
	context.SetParamValues(entities2.TestCarEntity.Vin.Vin)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities2.CarPersistenceEntity{
		"JH4DB1561NS000565": entities2.TestCarEntity,
		"JN8AZ2NC5B9300256": entities2.TestCarEntity,
		"2FDKF38G3KCA42390": entities2.TestCarEntity,
		"1GBJK39DX6E165432": entities2.TestCarEntity,
	})

	if assert.NoError(t, carsResource.GetCars(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, CarsBody, recorder.Body.String())
	}
}
