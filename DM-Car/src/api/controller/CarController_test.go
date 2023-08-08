package controller

import (
	"car/DM-Car/src/logic/model"
	"car/DM-Car/src/logic/operations"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockCarRepository struct {
	mockDatabase map[string]model.Car
}

func (mockRepository *MockCarRepository) Save(car model.Car) error {
	mockRepository.mockDatabase[car.Vin.Vin] = car
	return nil
}

func (mockRepository *MockCarRepository) FindAll() ([]model.Car, error) {
	cars := []model.Car{}
	for _, value := range mockRepository.mockDatabase {
		cars = append(cars, value)
	}
	return cars, nil
}

func (mockRepository *MockCarRepository) FindByVin(vin string) (model.Car, error) {
	car, ok := mockRepository.mockDatabase[vin]
	if ok {
		return car, nil
	}
	return model.Car{}, errors.New("Not found")
}

func TestPostCar(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}
	body := fmt.Sprintf(`
	{
		"vin": "%s",
		"brand": "%s",
		"model": "%s"
	}
	`, car.Vin.Vin, car.Brand, car.Model)
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/cars", strings.NewReader(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)

	carRepository := MockCarRepository{mockDatabase: map[string]model.Car{}}

	carOperations := operations.NewCarOperations(&carRepository)
	carsResource := NewCarController(carOperations)

	if assert.NoError(t, carsResource.PostCar(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, carRepository.mockDatabase, car.Vin.Vin)
		assert.Equal(t, carRepository.mockDatabase[car.Vin.Vin], car)
	}
}

func TestGetCarVin(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}
	uuid, _ := uuid.Parse(car.Vin.Vin)
	body := fmt.Sprintf(`{"Vin":{"Vin":"%s"},"Brand":"%s","Model":"%s"}
`, car.Vin.Vin, car.Brand, car.Model)
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/cars", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/:vin")
	context.SetParamNames("vin")
	context.SetParamValues(car.Vin.Vin)

	carRepository := MockCarRepository{mockDatabase: map[string]model.Car{
		"2af3d31e-15ef-11ee-be56-0242ac120005": car,
	}}

	carOperations := operations.NewCarOperations(&carRepository)
	carsResource := NewCarController(carOperations)

	if assert.NoError(t, carsResource.GetCarVin(context, uuid)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, body, recorder.Body.String())
	}
}

func TestGetCar(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}
	carBody := fmt.Sprintf(`{"Vin":{"Vin":"%s"},"Brand":"%s","Model":"%s"}`, car.Vin.Vin, car.Brand, car.Model)
	carsBody := fmt.Sprintf(`[%s,%s,%s,%s]
`, carBody, carBody, carBody, carBody)
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/cars", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/:vin")
	context.SetParamNames("vin")
	context.SetParamValues(car.Vin.Vin)

	carRepository := MockCarRepository{mockDatabase: map[string]model.Car{
		"2af3d31e-15ef-11ee-be56-0242ac120005": car,
		"2af3d31e-15ef-11ee-be56-0242ac120006": car,
		"2af3d31e-15ef-11ee-be56-0242ac120007": car,
		"2af3d31e-15ef-11ee-be56-0242ac120008": car,
	}}

	carOperations := operations.NewCarOperations(&carRepository)
	carsResource := NewCarController(carOperations)

	if assert.NoError(t, carsResource.GetCar(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, carsBody, recorder.Body.String())
	}
}
