package controller

import (
	"car/DM-Car/src/logic/model"
	"car/DM-Car/src/logic/operations"
	"car/DM-Car/src/support"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func CreateCarResourcesWithMockRepository(mockDatabaseContents map[string]model.Car) (CarController, operations.CarOperations, support.MockCarRepository) {
	carRepository := support.MockCarRepository{MockDatabase: mockDatabaseContents}
	carOperations := operations.NewCarOperations(&carRepository)
	return NewCarController(carOperations), carOperations, carRepository
}

// Test that 
func TestAddCar(t *testing.T) {
	context, request, recorder := support.CreateMockEchoSupport(
		http.MethodPost,
		"/cars",
		strings.NewReader(support.Body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, carRepository := CreateCarResourcesWithMockRepository(map[string]model.Car{})

	if assert.NoError(t, carsResource.AddCar(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, carRepository.MockDatabase, support.Car.Vin.Vin)
		assert.Equal(t, carRepository.MockDatabase[support.Car.Vin.Vin], support.Car)
	}
}

func TestAddCarWithExistingVin(t *testing.T) {
	context, request, _ := support.CreateMockEchoSupport(
		http.MethodPost,
		"/cars",
		strings.NewReader(support.Body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]model.Car{
		"JH4DA3350KS009715": support.Car,
	})

	assert.Error(t, carsResource.AddCar(context))
}

func TestAddCarInvalidVin(t *testing.T) {
	for _, invalidVin := range support.InvalidVins {
		body := fmt.Sprintf(`
		{
			"vin": "%s",
			"brand": "%s",
			"model": "%s"
		}
		`, invalidVin, support.Car.Brand, support.Car.Model)

		context, request, _ := support.CreateMockEchoSupport(
			http.MethodPost,
			"/cars",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]model.Car{})

		assert.Error(t, carsResource.AddCar(context))
	}
}

func TestAddCarNoBrand(t *testing.T) {
	body := fmt.Sprintf(`
	{
		"vin": "%s",
		"model": "%s"
	}
	`, support.Car.Vin.Vin, support.Car.Model)

	context, request, _ := support.CreateMockEchoSupport(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]model.Car{})

	assert.Error(t, carsResource.AddCar(context))
}

func TestAddCarNoModel(t *testing.T) {
	body := fmt.Sprintf(`
	{
		"vin": "%s",
		"brand": "%s"
	}
	`, support.Car.Vin.Vin, support.Car.Brand)

	context, request, _ := support.CreateMockEchoSupport(
		http.MethodPost,
		"/cars",
		strings.NewReader(body),
	)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]model.Car{})

	assert.Error(t, carsResource.AddCar(context))
}

func TestGetCar(t *testing.T) {
	body := fmt.Sprintf(`{"Vin":{"Vin":"%s"},"Brand":"%s","Model":"%s"}
`, support.Car.Vin.Vin, support.Car.Brand, support.Car.Model)

	context, _, recorder := support.CreateMockEchoSupport(
		http.MethodGet,
		"/cars",
		nil,
	)
	context.SetPath("/:vin")
	context.SetParamNames("vin")
	context.SetParamValues(support.Car.Vin.Vin)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]model.Car{
		"JH4DA3350KS009715": support.Car,
	})

	if assert.NoError(t, carsResource.GetCar(context, support.Car.Vin.Vin)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, body, recorder.Body.String())
	}
}

func TestGetCars(t *testing.T) {
	carBody := fmt.Sprintf(`{"Vin":{"Vin":"%s"},"Brand":"%s","Model":"%s"}`, support.Car.Vin.Vin, support.Car.Brand, support.Car.Model)
	carsBody := fmt.Sprintf(`[%s,%s,%s,%s]
`, carBody, carBody, carBody, carBody)

	context, _, recorder := support.CreateMockEchoSupport(
		http.MethodGet,
		"/cars",
		nil,
	)
	context.SetPath("/:vin")
	context.SetParamNames("vin")
	context.SetParamValues(support.Car.Vin.Vin)

	carsResource, _, _ := CreateCarResourcesWithMockRepository(map[string]model.Car{
		"JH4DA3350KS009715": support.Car,
		"JH4DA3350KS009716": support.Car,
		"JH4DA3350KS009717": support.Car,
		"JH4DA3350KS009718": support.Car,
	})

	if assert.NoError(t, carsResource.GetCars(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, carsBody, recorder.Body.String())
	}
}
