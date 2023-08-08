package operations

import (
	"car/DM-Car/src/logic/model"
	"car/DM-Car/src/support"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCars(t *testing.T) {
	carRepository := support.MockCarRepository{MockDatabase: map[string]model.Car{
		"JH4DA3350KS009715": support.Car,
	}}
	carOperations := NewCarOperations(&carRepository)

	cars, err := carOperations.GetCars()
	assert.Contains(t, cars, support.Car)
	assert.Equal(t, err, nil)
}
