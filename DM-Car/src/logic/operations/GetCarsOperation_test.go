package operations

import (
	"car/DM-Car/src/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCars(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}

	carRepository := MockCarRepository{mockDatabase: map[string]model.Car{
		"2af3d31e-15ef-11ee-be56-0242ac120005": car,
	}}
	carOperations := NewCarOperations(&carRepository)

	cars, err := carOperations.GetCars()
	assert.Contains(t, cars, car)
	assert.Equal(t, err, nil)
}
