package operations

import (
	"car/DM-Car/src/logic/model"
	"car/DM-Car/src/support"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of adding a car works
func TestAddCar(t *testing.T) {
	carRepository := support.MockCarRepository{MockDatabase: map[string]model.Car{}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.AddCar(support.Car.Vin.Vin, support.Car.Brand, support.Car.Model)
	assert.Equal(t, support.Car, carResult)
	assert.Equal(t, err, nil)
}
