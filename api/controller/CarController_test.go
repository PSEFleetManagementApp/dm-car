package controller

import (
	"car/infrastructure"
	"car/infrastructure/entities"
	"car/logic/operations"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// A valid request body for a Car
var CarBodyRequest = "{\"Vin\":\"JH4DA3350KS009715\",\"Brand\":\"Mercedes-Benz\",\"Model\":\"S Klasse\"}\n"

// A valid response body for a Car
var CarBodyResponse = "{\"Vin\":{\"Vin\":\"JH4DA3350KS009715\"},\"Brand\":\"Mercedes-Benz\",\"Model\":\"S Klasse\"}\n"

// A valid response body for Cars
var CarsBody = "{\"Cars\":[{\"Vin\":{\"Vin\":\"JH4DA3350KS009715\"},\"Brand\":\"Mercedes-Benz\",\"Model\":\"S Klasse\"},{\"Vin\":{\"Vin\":\"JH4DA3350KS009715\"},\"Brand\":\"Mercedes-Benz\",\"Model\":\"S Klasse\"},{\"Vin\":{\"Vin\":\"JH4DA3350KS009715\"},\"Brand\":\"Mercedes-Benz\",\"Model\":\"S Klasse\"},{\"Vin\":{\"Vin\":\"JH4DA3350KS009715\"},\"Brand\":\"Mercedes-Benz\",\"Model\":\"S Klasse\"}]}\n"

// List of invalid Vins according to the domain constraints
var InvalidVins = []string{
	"JH4DA3350KS00",
	"2CIGP44362R700796",
	"1C3CDZBG8DN5O4146",
	"1gCDC14K2LE198114",
	"1G3NF52E3XC4036521",
}

// Create all resources used by the car controller with an underlying mocked repository
func CreateCarResourcesWithMockRepository(mockDatabaseContents map[string]entities.CarPersistenceEntity) (CarController, operations.CarOperations, infrastructure.MockCarRepository) {
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

	carsResource, _, carRepository := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{})

	if assert.NoError(t, carsResource.AddCar(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, carRepository.MockDatabase, entities.TestCarEntity.Vin.Vin)
		assert.Equal(t, carRepository.MockDatabase[entities.TestCarEntity.Vin.Vin], entities.TestCarEntity)
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

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{
		"JH4DA3350KS009715": entities.TestCarEntity,
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
			entities.TestCarEntity.Brand,
			entities.TestCarEntity.Model)

		context, request, _ := CreateMockEcho(
			http.MethodPost,
			"/cars",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{})

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
		entities.TestCarEntity.Brand,
		entities.TestCarEntity.Model)

	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{})

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
		entities.TestCarEntity.Vin.Vin,
		entities.TestCarEntity.Model)

	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{})

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
		entities.TestCarEntity.Vin.Vin,
		entities.TestCarEntity.Brand)

	context, request, _ := CreateMockEcho(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{})

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
	context.SetParamValues(entities.TestCarEntity.Vin.Vin)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{
		"JH4DA3350KS009715": entities.TestCarEntity,
	})

	if assert.NoError(t, carsResource.GetCar(context, entities.TestCarEntity.Vin.Vin)) {
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
	context.SetParamValues(entities.TestCarEntity.Vin.Vin)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]entities.CarPersistenceEntity{
		"JH4DA3350KS009715": entities.TestCarEntity,
		"JH4DA3350KS009716": entities.TestCarEntity,
		"JH4DA3350KS009717": entities.TestCarEntity,
		"JH4DA3350KS009718": entities.TestCarEntity,
	})

	if assert.NoError(t, carsResource.GetCars(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, CarsBody, recorder.Body.String())
	}
}
