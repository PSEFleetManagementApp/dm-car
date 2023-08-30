package operations

import (
	"car/infrastructure"
	"car/infrastructure/persistenceentities"
	"car/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of adding a car works
func TestAddCar(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]persistenceentities.CarPersistenceEntity{}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.AddCar(model.TestCarModel.Vin.Vin, model.TestCarModel.Brand, model.TestCarModel.Model)
	assert.Equal(t, model.TestCarModel, carResult)
	assert.Equal(t, err, nil)
}
