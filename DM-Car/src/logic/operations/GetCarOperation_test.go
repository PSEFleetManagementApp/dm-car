package operations

import (
	"car/DM-Car/src/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCar(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}

	carRepository := MockCarRepository{mockDatabase: map[string]model.Car{
		"2af3d31e-15ef-11ee-be56-0242ac120005": car,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar(car.Vin.Vin)
	assert.Equal(t, car, carResult)
	assert.Equal(t, err, nil)
}
