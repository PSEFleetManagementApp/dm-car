package operations

import (
	"car/DM-Car/src/logic/model"
	"car/DM-Car/src/support"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCar(t *testing.T) {
	carRepository := support.MockCarRepository{MockDatabase: map[string]model.Car{
		"JH4DA3350KS009715": support.Car,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar(support.Car.Vin.Vin)
	assert.Equal(t, support.Car, carResult)
	assert.Equal(t, err, nil)
}
